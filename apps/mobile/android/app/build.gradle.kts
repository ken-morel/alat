plugins {
    id("com.android.application")
    id("kotlin-android")
    // The Flutter Gradle Plugin must be applied after the Android and Kotlin Gradle plugins.
    id("dev.flutter.flutter-gradle-plugin")
}

android {
    namespace = "cm.engon.alat.mobile"
    compileSdk = 36

    defaultConfig {
        // TODO: Specify your own unique Application ID (https://developer.android.com/studio/build/application-id).
        applicationId = "com.example.alat"
        // You can update the following values to match your application needs.
        // For more information, see: https://docs.flutter.dev/deployment/android#reviewing-the-build-configuration.
        minSdk = flutter.minSdkVersion
        targetSdk = 34
        versionCode = flutter.versionCode
        versionName = flutter.versionName
    }

    compileOptions {
        isCoreLibraryDesugaringEnabled = true
    }


    kotlin {
        jvmToolchain(8)
    }

    buildTypes {
        release {
            // TODO: Add your own signing config for the release build.
            // See https://developer.android.com/studio/publish/app-signing#secure-your-key for more information.
            signingConfig = signingConfigs.getByName("debug")
        }
    }
}

flutter {
    source = "../.."
}

dependencies {
    coreLibraryDesugaring("com.android.tools:desugar_jdk_libs:2.1.4")
}
