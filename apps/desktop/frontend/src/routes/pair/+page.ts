import { SearchDevices } from "$lib/wails/wailsjs/go/app/App";
import type { PageLoad } from "./$types";

export const load: PageLoad = async () => {
  await SearchDevices();
  return {};
};
