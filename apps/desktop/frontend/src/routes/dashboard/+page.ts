import { GetPairedDevices } from "$lib/wails/wailsjs/go/app/App";
import type { PageLoad } from "./$types";

export const load: PageLoad = async ({}) => {
  const paired = await GetPairedDevices();
  return {
    paired: paired ? paired : [],
  };
};
