import 'dart:async';

import 'package:flutter/foundation.dart';
import 'package:flutter/widgets.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:intl/intl.dart' as intl;

import 'app_localizations_en.dart';
import 'app_localizations_fr.dart';

// ignore_for_file: type=lint

/// Callers can lookup localized strings with an instance of AppLocalizations
/// returned by `AppLocalizations.of(context)`.
///
/// Applications need to include `AppLocalizations.delegate()` in their app's
/// `localizationDelegates` list, and the locales they support in the app's
/// `supportedLocales` list. For example:
///
/// ```dart
/// import 'l10n/app_localizations.dart';
///
/// return MaterialApp(
///   localizationsDelegates: AppLocalizations.localizationsDelegates,
///   supportedLocales: AppLocalizations.supportedLocales,
///   home: MyApplicationHome(),
/// );
/// ```
///
/// ## Update pubspec.yaml
///
/// Please make sure to update your pubspec.yaml to include the following
/// packages:
///
/// ```yaml
/// dependencies:
///   # Internationalization support.
///   flutter_localizations:
///     sdk: flutter
///   intl: any # Use the pinned version from flutter_localizations
///
///   # Rest of dependencies
/// ```
///
/// ## iOS Applications
///
/// iOS applications define key application metadata, including supported
/// locales, in an Info.plist file that is built into the application bundle.
/// To configure the locales supported by your app, you’ll need to edit this
/// file.
///
/// First, open your project’s ios/Runner.xcworkspace Xcode workspace file.
/// Then, in the Project Navigator, open the Info.plist file under the Runner
/// project’s Runner folder.
///
/// Next, select the Information Property List item, select Add Item from the
/// Editor menu, then select Localizations from the pop-up menu.
///
/// Select and expand the newly-created Localizations item then, for each
/// locale your application supports, add a new item and select the locale
/// you wish to add from the pop-up menu in the Value field. This list should
/// be consistent with the languages listed in the AppLocalizations.supportedLocales
/// property.
abstract class AppLocalizations {
  AppLocalizations(String locale)
    : localeName = intl.Intl.canonicalizedLocale(locale.toString());

  final String localeName;

  static AppLocalizations? of(BuildContext context) {
    return Localizations.of<AppLocalizations>(context, AppLocalizations);
  }

  static const LocalizationsDelegate<AppLocalizations> delegate =
      _AppLocalizationsDelegate();

  /// A list of this localizations delegate along with the default localizations
  /// delegates.
  ///
  /// Returns a list of localizations delegates containing this delegate along with
  /// GlobalMaterialLocalizations.delegate, GlobalCupertinoLocalizations.delegate,
  /// and GlobalWidgetsLocalizations.delegate.
  ///
  /// Additional delegates can be added by appending to this list in
  /// MaterialApp. This list does not have to be used at all if a custom list
  /// of delegates is preferred or required.
  static const List<LocalizationsDelegate<dynamic>> localizationsDelegates =
      <LocalizationsDelegate<dynamic>>[
        delegate,
        GlobalMaterialLocalizations.delegate,
        GlobalCupertinoLocalizations.delegate,
        GlobalWidgetsLocalizations.delegate,
      ];

  /// A list of this localizations delegate's supported locales.
  static const List<Locale> supportedLocales = <Locale>[
    Locale('en'),
    Locale('fr'),
  ];

  /// No description provided for @aShortMemorableName.
  ///
  /// In en, this message translates to:
  /// **'A short, memorable name'**
  String get aShortMemorableName;

  /// No description provided for @permitConnectedDevicesToSendFilesToThisDevicesWithouthPriorConfirmation.
  ///
  /// In en, this message translates to:
  /// **'Permit connected devices to send files to this devices, withouth prior confirmation'**
  String
  get permitConnectedDevicesToSendFilesToThisDevicesWithouthPriorConfirmation;

  /// No description provided for @helloWorld.
  ///
  /// In en, this message translates to:
  /// **'Hello World!'**
  String get helloWorld;

  /// No description provided for @services.
  ///
  /// In en, this message translates to:
  /// **'Services'**
  String get services;

  /// No description provided for @weAreGoingToGoThroughOutTheProcessOfSettingUpYourDevice.
  ///
  /// In en, this message translates to:
  /// **'We are going to go through out the process of setting up your device'**
  String get weAreGoingToGoThroughOutTheProcessOfSettingUpYourDevice;

  /// No description provided for @serviceDisabled.
  ///
  /// In en, this message translates to:
  /// **'Service Disabled'**
  String get serviceDisabled;

  /// No description provided for @becauseEachOfYourDevicesAlsoMeritsAName.
  ///
  /// In en, this message translates to:
  /// **'Because each of your devices also merits a name'**
  String get becauseEachOfYourDevicesAlsoMeritsAName;

  /// No description provided for @yourDeviceIsCompletelySetupAndNowReadyToStart.
  ///
  /// In en, this message translates to:
  /// **'Your device is completely setup and now ready to start'**
  String get yourDeviceIsCompletelySetupAndNowReadyToStart;

  /// No description provided for @errorInitializingAlatCoreError.
  ///
  /// In en, this message translates to:
  /// **'Error initializing alat core: {error}'**
  String errorInitializingAlatCoreError(Object error);

  /// No description provided for @theSettingsAreAlreadyInSensibleDefaultsWhichShouldJustFitYouButReadAttentivelySoYouKnowWhatYouExposeToOthers.
  ///
  /// In en, this message translates to:
  /// **'The settings are already in sensible defaults which should just fit you, but read attentively so you know what you expose to others.'**
  String
  get theSettingsAreAlreadyInSensibleDefaultsWhichShouldJustFitYouButReadAttentivelySoYouKnowWhatYouExposeToOthers;

  /// No description provided for @name.
  ///
  /// In en, this message translates to:
  /// **'Name'**
  String get name;

  /// No description provided for @fileReceive.
  ///
  /// In en, this message translates to:
  /// **'File receive'**
  String get fileReceive;

  /// No description provided for @previous.
  ///
  /// In en, this message translates to:
  /// **'Previous'**
  String get previous;

  /// No description provided for @alatSetup.
  ///
  /// In en, this message translates to:
  /// **'Alat setup'**
  String get alatSetup;

  /// No description provided for @serviceEnabled.
  ///
  /// In en, this message translates to:
  /// **'Service Enabled'**
  String get serviceEnabled;

  /// No description provided for @initializingAlatCore.
  ///
  /// In en, this message translates to:
  /// **'Initializing alat core...'**
  String get initializingAlatCore;

  /// No description provided for @thisPermitsOtherDevicesToDisplayThisDevicesBatteryMemoryAndOtherSystemInformation.
  ///
  /// In en, this message translates to:
  /// **'This permits other devices to display this device\'s battery, memory, and other system information.'**
  String
  get thisPermitsOtherDevicesToDisplayThisDevicesBatteryMemoryAndOtherSystemInformation;

  /// No description provided for @nameMe.
  ///
  /// In en, this message translates to:
  /// **'Name me'**
  String get nameMe;

  /// No description provided for @next.
  ///
  /// In en, this message translates to:
  /// **'Next'**
  String get next;

  /// No description provided for @aColorToMoreEasilyIdentifyYourDevice.
  ///
  /// In en, this message translates to:
  /// **'A color to more easily identify your device'**
  String get aColorToMoreEasilyIdentifyYourDevice;

  /// No description provided for @servicesAreTheDifferentFeaturesOfYourDeviceYouWantToMakeAvailableToConnectedDevicesYouCanDisableThemAtAnyTime.
  ///
  /// In en, this message translates to:
  /// **'Services are the different features of your device you want to make available to connected devices. You can disable them at any time.'**
  String
  get servicesAreTheDifferentFeaturesOfYourDeviceYouWantToMakeAvailableToConnectedDevicesYouCanDisableThemAtAnyTime;

  /// No description provided for @systemInformationAndStats.
  ///
  /// In en, this message translates to:
  /// **'System Information and Stats'**
  String get systemInformationAndStats;

  /// No description provided for @seconds.
  ///
  /// In en, this message translates to:
  /// **'seconds'**
  String get seconds;

  /// No description provided for @megaBytes.
  ///
  /// In en, this message translates to:
  /// **'Mega bytes'**
  String get megaBytes;

  /// No description provided for @options.
  ///
  /// In en, this message translates to:
  /// **'Options'**
  String get options;

  /// No description provided for @initializingAlat.
  ///
  /// In en, this message translates to:
  /// **'Initializing Alat...'**
  String get initializingAlat;

  /// No description provided for @maximulFileSize.
  ///
  /// In en, this message translates to:
  /// **'Maximul file size'**
  String get maximulFileSize;

  /// No description provided for @setupComplete.
  ///
  /// In en, this message translates to:
  /// **'Setup complete'**
  String get setupComplete;

  /// No description provided for @cacheRefreshInterval.
  ///
  /// In en, this message translates to:
  /// **'Cache Refresh Interval'**
  String get cacheRefreshInterval;

  /// No description provided for @deviceColor.
  ///
  /// In en, this message translates to:
  /// **'Device color'**
  String get deviceColor;

  /// No description provided for @deviceName.
  ///
  /// In en, this message translates to:
  /// **'Device name'**
  String get deviceName;

  /// No description provided for @welcomeToAlat.
  ///
  /// In en, this message translates to:
  /// **'Welcome to alat'**
  String get welcomeToAlat;
}

class _AppLocalizationsDelegate
    extends LocalizationsDelegate<AppLocalizations> {
  const _AppLocalizationsDelegate();

  @override
  Future<AppLocalizations> load(Locale locale) {
    return SynchronousFuture<AppLocalizations>(lookupAppLocalizations(locale));
  }

  @override
  bool isSupported(Locale locale) =>
      <String>['en', 'fr'].contains(locale.languageCode);

  @override
  bool shouldReload(_AppLocalizationsDelegate old) => false;
}

AppLocalizations lookupAppLocalizations(Locale locale) {
  // Lookup logic when only language code is specified.
  switch (locale.languageCode) {
    case 'en':
      return AppLocalizationsEn();
    case 'fr':
      return AppLocalizationsFr();
  }

  throw FlutterError(
    'AppLocalizations.delegate failed to load unsupported locale "$locale". This is likely '
    'an issue with the localizations generation tool. Please file an issue '
    'on GitHub with a reproducible sample app and the gen-l10n configuration '
    'that was used.',
  );
}
