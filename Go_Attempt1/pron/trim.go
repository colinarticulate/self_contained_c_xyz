package pron

import (
	"io"
	"log"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/cryptix/wav"
	"github.com/maxhawkins/go-webrtcvad"
)

type voicedInterval struct {
	start, duration int
}

//===================================================================================================
//  __      __   _    ___ _____ ___    __   ___   ___              _   _   _
//  \ \    / /__| |__| _ \_   _/ __|   \ \ / /_\ |   \     ___ ___| |_| |_(_)_ _  __ _ ___
//   \ \/\/ / -_) '_ \   / | || (__     \ V / _ \| |) |   (_-</ -_)  _|  _| | ' \/ _` (_-<
//    \_/\_/\___|_.__/_|_\ |_| \___|     \_/_/ \_\___/    /__/\___|\__|\__|_|_||_\__, /__/
//                                                                               |___/
//===================================================================================================
// The number used with WebRtc sets the VAD operating mode. A more aggressive (higher mode) VAD is more
// restrictive in reporting speech. Put in other words the probability of being speech when the VAD
// returns 1 is increased with increasing mode. As a consequence also the missed detection rate goes up.
//
// Aggressiveness mode (0, 1, 2, or 3).
/*
VAD.Mode.NORMAL   ... 0?
Constant for normal voice detection mode. Suitable for high bitrate, low-noise data. May classify noise as voice, too. The default value if mode is omitted in the constructor.

VAD.Mode.LOW_BITRATE   ... 1?
Detection mode optimised for low-bitrate audio.

VAD.Mode.AGGRESSIVE   ..... 2?
Detection mode best suited for somewhat noisy, lower quality audio.

VAD.Mode.VERY_AGGRESSIVE  ..... 3?
Detection mode with lowest miss-rate. Works well for most inputs.


The WebRTC VAD only accepts 16-bit mono PCM audio, sampled at 8000, 16000, 32000 or 48000 Hz. A frame must be either 10, 20, or 30 ms in duration:
Optionally, set its aggressiveness mode, which is an integer between 0 and 3. 0 is the least aggressive about filtering out non-speech, 3 is the most aggressive.
*/

/*
Addtionaly
==========
The VAD engine requires mono, 16-bit PCM audio with a sample rate of 8, 16, 32 or 48 KHz as input. The input should be an audio segment of 10, 20 or 30 milliseconds.
When the audio input is 16 Khz, the input array should thus be either of length 160, 320 or 480.
https://github.com/jitsi/jitsi-webrtc-vad-wrapper

For example, if your sample rate is 16000 Hz, then the only allowed frame/chunk sizes are 16000 * ({10,20,30} / 1000) = 160, 320 or 480 samples.
Since each sample is 2 bytes (16 bits), the only allowed frame/chunk sizes are 320, 640, or 960 bytes.
https://github.com/wiseman/py-webrtcvad/issues/30
*/

func fetchTrimBounds(audiofile string, proffile string, phons []phoneme) (float64, float64) {
	// Original we set this quite aggressively, but may?? be causing issues when background is noisy, needs more investigation - PE August 2019
	//start2, duration2 := webRtcBounds(audiofile, 2)
	//start3, _ := webRtcBounds(audiofile, 3)
	dir, file := filepath.Split(audiofile)
	ext := filepath.Ext(file)
	noise_profile_removed := filepath.Join(dir, file[:len(file)-len(ext)]+"_noise_profile_removed_"+ext)
	beep_removed := filepath.Join(dir, file[:len(file)-len(ext)]+"_beep_removed_"+ext)
	normalised := filepath.Join(dir, file[:len(file)-len(ext)]+"_normalised_"+ext)
	normalised2 := filepath.Join(dir, file[:len(file)-len(ext)]+"_normalised2_"+ext)
	sinc_applied := filepath.Join(dir, file[:len(file)-len(ext)]+"_sinc_applied_"+ext)
	highpass := filepath.Join(dir, file[:len(file)-len(ext)]+"_highpass_"+ext)
	tone_removed := filepath.Join(dir, file[:len(file)-len(ext)]+"_tone_removed_"+ext)
	double_tone_removed := filepath.Join(dir, file[:len(file)-len(ext)]+"_double_tone_removed_"+ext)

	out, err := exec.Command("soxi", "-D", audiofile).Output()
	if err != nil {
		log.Panic(err)
	}
	// Yuk! out contains a float terminated by '\n' so strip the '\n'. There must
	// be a better way...
	//
	length_audio, err := strconv.ParseFloat(string(out[:len(out)-1]), 64)
	if err != nil {
		log.Panic(err)
	}

	//*******************************
	// Using
	// High pass (250Hz) --> band reject (650-550Hz) to remove the beep
	// then a second time --> band reject (650-550Hz) to really push down the signal
	// create a noise profile (currently applied to whole audio, not sure if this is best way?)
	// remove that noise profile from the audio
	// Normalise
	// remove beep profile
	// 2nd noramlise
	// Apply sinc function (reject everything below 300Hz and above 4500Hz) ... see notes below

	_, err = exec.Command("sox", audiofile, highpass, "sinc", "250-0.1").Output()
	if err != nil {
		debug("Step 1: High pass filter failed ", err)
	}

	_, err = exec.Command("sox", highpass, tone_removed, "sinc", "650-550").Output()
	if err != nil {
		debug("Step 2: Tone removal failed ", err)
	}

	_, err = exec.Command("sox", tone_removed, double_tone_removed, "sinc", "650-550").Output()
	if err != nil {
		debug("Step 3: Tone removal failed ", err)
	}

	new_name := audiofile + "_noise.prof"
	_, err = exec.Command("sox", double_tone_removed, "-n", "noiseprof", new_name).Output()
	if err != nil {
		debug("Step 4: Can't create noise profile, ", err)
	}
	_, err = exec.Command("sox", double_tone_removed, noise_profile_removed, "noisered", new_name, "0.21").Output()
	if err != nil {
		debug("Step 5: Can't apply noise profile, ", err)
	}

	_, err = exec.Command("sox", noise_profile_removed, normalised, "norm", "-0.1").Output()
	if err != nil {
		debug("Step 6: Normalisation failed ", err)
	}

	// beep_profile2 := "/Users/test/Documents/GitHub/test_pronounce/test_script/beep_noise2.prof"
	_, err = exec.Command("sox", normalised, beep_removed, "noisered", proffile, "0.05").Output()
	if err != nil {
		debug("Step 7: Can't remove the beep, ", err)
	}

	_, err = exec.Command("sox", beep_removed, normalised2, "norm", "-0.1").Output()
	if err != nil {
		debug("Step 8: 2nd Normalisation failed ", err)
	}

	// Not applying ang agc because it makes the 'beep' and 'audio' the same amplitude (ie put energy back into the beep)
	// The webRTC vad stage then thinks the beep is the desird audio??
	//_, err = exec.Command("sox", normalised2, agc_ed, "compand", "0.2,0.2", "-40,-40,-35,-20,0,-20", "-10", "-60", "0.1").Output()
	//if err != nil {
	//debug("AGC failed, ", err)
	//debug("sox", audiofile, agc_ed, "compand", "0.2,0.2", "-40,-40,-35,-20,0,-20", "-10", "-60", "0.1")
	// }

	// Sinc filter
	//Apply a sinc kaiser-windowed low-pass, high-pass, band-pass, or band-reject filter to the signal.
	//The freqHP and freqLP parameters give the frequencies of the 6dB points of a high-pass and low-pass filter that may be invoked individually
	//or together. If both are given, then freqHP less than freqLP creates a band-pass filter
	//In telephony, the usable voice frequency band ranges from approximately 300 to 3400 Hz

	_, err = exec.Command("sox", normalised2, sinc_applied, "sinc", "300-4500").Output()

	start2 := 0.0
	duration2 := 0.0
	start3 := 0.0
	duration3 := 0.0

	start2, duration2 = webRtcBounds(sinc_applied, 1)
	start3, duration3 = webRtcBounds(sinc_applied, 3)

	start := 0.0
	end := 0.0
	duration := 0.0

	end2 := start2 + duration2
	end3 := start3 + duration3
	end = math.Max(end2, end3)

	debug("\n webRTC stage .... start2, duration2, end2 = ", start2, duration2, end2)
	debug(" webRTC stage .... start3, duration3, end3 = ", start3, duration3, end3)

	//new 18 Oct 2022
	// Trying to deal with when a user repeats the word over and over

	if start2 >= (2 * start3) {
		start3 = start2
		end = math.Max(end2, end3)
	}
	if start3 >= (2 * start2) {
		start2 = start3
		end = math.Max(end2, end3)
		//end = end2;
	}

	debug("\n Align starts/end if repeat audio .... start2, duration2, end2 = ", start2, duration2, end2)
	debug(" Align starts/end if repeat audio .... start3, duration3, end3 = ", start3, duration3, end3)

	start = math.Min(start2, start3)

	debug(" Initial choice for start = ", start)

	// Use the settings below for production - PE 2 June 2022 -- the ones above are for trimming scraped audio for training
	if start-0.2 <= 0 {
		start = 0
	} else {
		start = start - 0.10
		// This is required for the S words where VAD starts late, example seat1_tressa, the s goes missing without an earlier 0 period
	}
	debug(" Reduce position of start (if not too close to 0) = ", start)

	duration = end - start

	debug("\n Initial start, duration, end = ", start, duration, end)

	//Adjustments
	//start = start - 0.2333  //*********************** adjusted to try and fix tomo replied where the voiceless plosive "p" goes to zero and VAD thinks it's a silence gap
	// 0.4 is added to cover contraceptives1_tomo  .... and potentially others than have voiceless plosives towards the end ... though big silent gaps could be an issue

	duration = duration + 0.45

	//check if start + duration is greater than the end of the audio file
	if (start + duration) >= length_audio {
		duration = length_audio - start - 0.02
	}

	guard_end := 0.0
	guard_end = start + duration

	debug("\n Cut at: start, duration = ", start, duration, "    (guard_end = ", guard_end, ")\n")

	return start, duration

}

func webRtcBounds(audiofile string, mode int) (float64, float64) {
	info, err := os.Stat(audiofile)
	if err != nil {
		debug("webRtcBounds: call to os.Stat failed. err =", err)
		log.Panic(err)
	}

	file, err := os.Open(audiofile)
	if err != nil {
		debug("webRtcBounds: failed to open file. err =", err)
		log.Panic(err)
	}
	defer file.Close()

	wavReader, err := wav.NewReader(file, info.Size())
	if err != nil {
		debug("webRtcBounds: call to wav.NewReader failed. err =", err)
		log.Panic(err)
	}

	reader, err := wavReader.GetDumbReader()
	if err != nil {
		debug("webRtcBounds: call to wav.GetDumbReader failed. err =", err)
		log.Panic(err)
	}

	wavInfo := wavReader.GetFile()
	rate := int(wavInfo.SampleRate)
	if wavInfo.Channels != 1 {
		debug("webRtcBounds: expected mono file")
		log.Panic("expected mono file")
	}
	if rate != 16000 {
		debug("webRtcBounds: expected 16kHz file")
		log.Panic("expected 16kHz file")
	}

	vad, err := webrtcvad.New()
	// vad, err := artVad.New()
	if err != nil {
		debug("webRtcBounds: call to webrtcvad.New failed. err =", err)
		log.Panic(err)
	}

	if err := vad.SetMode(mode); err != nil {
		debug("webRtcBounds: call to vad.SetMode failed. err =", err)
		log.Panic(err)
	}

	//frame := make([]byte, 160*2)               //  160/16kHz=10ms but each sample is 2 bytes.  Therefore, 320=10ms, 640=20ms, 960=30ms.
	frame := make([]byte, 640) //640

	debug("\n\n\nlen(frame)=", len(frame))
	debug("rate =", rate)
	//debug("\nThe frame is = ", frame)
	//if ok := vad.ValidRateAndFrameLength(rate, len(frame)); !ok {
	//  debug("\nwebRtcBounds: Valid frame rate & length checker failed")
	//  log.Panic("webRTC: invalid rate or frame length")
	//}

	var isActive bool
	var offset int

	var start, duration int
	var thisStart int

	report := func() {
		_ = time.Duration(offset) * time.Second / time.Duration(rate) / 2
	}

	for {
		_, err := io.ReadFull(reader, frame)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		}
		if err != nil {
			debug("webRtcBounds: call to io.ReadFull failed. err =", err)
			log.Panic(err)
		}

		frameActive, err := vad.Process(rate, frame)
		if err != nil {
			debug("webRtcBounds: call to vad.Process failed. err =", err)
			log.Panic(err)
		}

		if isActive != frameActive || offset == 0 {
			// isActive = frameActive
			report()
		}
		// Find word on the fly
		// Declare start, duration, both initialised to 0
		// On isActive set potStart, on !isactive check to see if duration > last
		// saved duration. If it is start = potStart, duration == 'new duration'
		if isActive != frameActive || offset == 0 {
			if isActive {
				if offset-thisStart > duration {
					start = thisStart
					duration = offset - start
				}
			} else {
				thisStart = offset
			}
			isActive = frameActive
		}
		offset += len(frame)
	}
	report()
	if isActive {
		if offset-thisStart > duration {
			start = thisStart
			duration = offset - start
		}
	}
	startSecs := float64(start) / float64(rate*2)
	if startSecs > 0.1 {
		startSecs -= 0.1
	}
	durationSecs := (float64(duration) / float64(rate*2)) + 0.1
	return startSecs, durationSecs
}

func trimAudio(audiofile string, proffile string, phons []phoneme) string {
	start, duration := fetchTrimBounds(audiofile, proffile, phons)

	dir, file := filepath.Split(audiofile)
	ext := filepath.Ext(file)
	trimmedfile := filepath.Join(dir, file[:len(file)-len(ext)]+"_trimmed"+ext)

	cmd := exec.Command("sox", audiofile, trimmedfile, "trim", strconv.FormatFloat(start, 'f', -1, 64), strconv.FormatFloat(duration, 'f', -1, 64))

	err := cmd.Run()
	if err != nil {
		debug(" ")
		debug("Command call to sox failed. err =", err)
		debug("exit status 2 (if that's it) can mean a missing parameter")
		debug("command sent was:-")
		debug("sox", audiofile, trimmedfile, "trim", strconv.FormatFloat(start, 'f', -1, 64), strconv.FormatFloat(duration, 'f', -1, 64))
		debug("Should be:-")
		debug("sox inputFile outputFile trim start duration")
		debug(" ")

	}
	return trimmedfile
}
