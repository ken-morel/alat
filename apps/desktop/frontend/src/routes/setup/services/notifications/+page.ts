import { SettingsGetNotificationSyncSettings } from "$lib/wails/wailsjs/go/app/App";
import type { config } from "$lib/wails/wailsjs/go/models";

export const load = async (): Promise<{
  settings: config.NotificationSyncSettings;
}> => {
  const settings = await SettingsGetNotificationSyncSettings();
  return {
    settings,
  };
};
