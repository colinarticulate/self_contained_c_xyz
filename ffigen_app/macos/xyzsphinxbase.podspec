#
# To learn more about a Podspec see http://guides.cocoapods.org/syntax/podspec.html.
# Run `pod lib lint ffigen_app.podspec` to validate before publishing.
#
Pod::Spec.new do |s|
    s.name             = 'xyzsphinxbase'
    s.version          = '0.0.1'
    s.summary          = 'Sphinx Base module'
    s.description      = <<-DESC
  A new Flutter FFI plugin project.
                         DESC
    s.homepage         = 'http://example.com'
    s.license          = { :file => '../LICENSE' }
    s.author           = { 'Your Company' => 'email@example.com' }
  
    # This will ensure the source files in Classes/ are included in the native
    # builds of apps using this FFI plugin. Podspec does not support relative
    # paths, so Classes contains a forwarder C file that relatively imports
    # `../src/*` so that the C sources can be shared among all target platforms.
    s.source           = { :path => '.' }
    s.source_files     = ['Classes/src/xyzsphinxbase/src/libsphinxbase/fe/*',
                            'Classes/src/xyzsphinxbase/src/libsphinxbase/feat/*',
                            'Classes/src/xyzsphinxbase/src/libsphinxbase/lm/*',
                            'Classes/src/xyzsphinxbase/src/libsphinxbase/util/*']
    s.public_header_files = ['Classes/src/xyzsphinxbase/include/xyzsphinxbase/*']
    s.private_header_files = ['Classes/src/xyzsphinxbase/include/*.h']
    #s.dependency 'FlutterMacOS'
  
    s.platform = :osx, '10.11'
    s.pod_target_xcconfig = { 'DEFINES_MODULE' => 'YES' }
    s.swift_version = '5.0'
    s.libraries = 'stdc'
    s.mdoule_name ='xyzsphinxbase'
    #s.dependency 'ffigen_app'
   
  end
  