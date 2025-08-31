import { SettingsGetRemoteInputSettings } from "$lib/wails/wailsjs/go/app/App";
import type { config } from "$lib/wails/wailsjs/go/models";

export const load = async (): Promise<{
  settings: config.RemoteInputSettings;
}> => {
  const settings = await SettingsGetRemoteInputSettings();
  return {
    settings,
  };
};
