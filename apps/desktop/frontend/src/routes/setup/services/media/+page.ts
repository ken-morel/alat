import { SettingsGetMediaControlSettings } from "$lib/wails/wailsjs/go/app/App";
import type { config } from "$lib/wails/wailsjs/go/models";

export const load = async (): Promise<{
  settings: config.MediaControlSettings;
}> => {
  const settings = await SettingsGetMediaControlSettings();
  return {
    settings,
  };
};
