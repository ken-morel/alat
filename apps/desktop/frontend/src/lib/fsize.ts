export const B = 1;
export const KB = B * 1024;
export const MB = KB * 1024;
export const GB = MB * 1024;
export const TB = GB * 1024;

export type LabelType = "short" | "long";
type LabelGroup = [b: string, k: string, m: string, g: string, t: string];

const LABELS: Record<LabelType, LabelGroup> = {
  short: ["B", "kB", "MB", "GB", "TB"],
  long: ["Bytes", "Kilobytes", "Megabytes", "Gigabytes", "Terabytes"],
};

export default function fsizeText(
  size: number,
  options: {
    decimals: number;
    label: LabelType;
    labels?: LabelGroup;
  } = {
    decimals: 1,
    label: "short",
  },
): string {
  const lbl: LabelGroup = options.labels
    ? options.labels
    : LABELS[options.label];
  if (size > TB) {
    return (size / TB).toFixed(options.decimals).toString() + lbl[4];
  } else if (size > GB) {
    return (size / GB).toFixed(options.decimals).toString() + lbl[3];
  } else if (size > MB) {
    return (size / MB).toFixed(options.decimals).toString() + lbl[2];
  } else if (size > KB) {
    return (size / KB).toFixed(options.decimals).toString() + lbl[1];
  } else {
    return size.toFixed(options.decimals).toString() + lbl[0];
  }
}
