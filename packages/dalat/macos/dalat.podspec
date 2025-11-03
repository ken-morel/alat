#
# To learn more about a Podspec see http://guides.cocoapods.org/syntax/podspec.html.
# Run `pod lib lint dalat.podspec` to validate before publishing.
#
Pod::Spec.new do |s|
  s.name             = 'dalat'
  s.version          = '0.1.0'
  s.summary          = 'Alat library dart plugin'
  s.description      = <<-DESC
Alat library dart plugin
                       DESC
  s.homepage         = 'https://github.com/ken-morel/alat'
  s.license          = { :file => '../LICENSE' }
  s.author           = { 'Engon Ken Morel' => 'engonken8@gmail.com' }

  # This will ensure the source files in Classes/ are included in the native
  # builds of apps using this FFI plugin. Podspec does not support relative
  # paths, so Classes contains a forwarder C file that relatively imports
  # `../src/*` so that the C sources can be shared among all target platforms.
  s.source           = { :path => '.' }
  s.source_files = 'Classes/**/*'

  # If your plugin requires a privacy manifest, for example if it collects user
  # data, update the PrivacyInfo.xcprivacy file to describe your plugin's
  # privacy impact, and then uncomment this line. For more information,
  # see https://developer.apple.com/documentation/bundleresources/privacy_manifest_files
  # s.resource_bundles = {'dalat_privacy' => ['Resources/PrivacyInfo.xcprivacy']}

  s.dependency 'FlutterMacOS'

  s.platform = :osx, '10.11'
  s.pod_target_xcconfig = { 'DEFINES_MODULE' => 'YES' }
  s.swift_version = '5.0'
end
