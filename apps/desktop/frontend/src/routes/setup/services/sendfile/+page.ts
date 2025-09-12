import { SettingsGetFileSend } from "$lib/wails/wailsjs/go/app/App";

export const load = async () => {
  const settings = await SettingsGetFileSend();
  return {
    settings,
  };
};
