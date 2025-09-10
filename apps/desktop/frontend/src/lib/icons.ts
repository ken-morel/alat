import {
  Tv,
  Smartphone,
  Laptop,
  LucideGhost,
  Cpu,
  Network,
} from "lucide-svelte";

import type { DeviceType } from "./device";

export const iconMap: Record<DeviceType, typeof Tv> = {
  tv: Tv,
  mobile: Smartphone,
  desktop: Laptop,
  arduino: Cpu,
  web: Network,
  unspecified: LucideGhost,
};

export const Ghost = LucideGhost;
export default (t: string) => iconMap[t as DeviceType] || Ghost;
