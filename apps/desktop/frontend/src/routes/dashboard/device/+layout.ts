import { connectedDevice } from "$lib/store";
import { get } from "svelte/store";
import { goto } from "$app/navigation";

export const load = () => {
  try {
    if (!get(connectedDevice)) goto("/dashboard");
  } catch (e) {}
  return {};
};
