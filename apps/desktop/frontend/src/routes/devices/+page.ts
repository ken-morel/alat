import type { PageLoad } from "./$types";
import { GetFoundDevices } from "$lib/wails/wailsjs/go/app/App";

export const load: PageLoad = async () => {
  return {
    found: (await GetFoundDevices()) || [],
  };
};
