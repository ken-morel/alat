/* w3color.js ver.1.11 by w3schools.com (Do not remove this line)*/
/*
camelcase(name::String) = lowercase(name[1]) * name[2:end]
function snakecase(name::String) 
  name = name |> camelcase
  while (idx = findfirst(isuppercase, name); !isnothing(idx))
    name = name[1:idx-1] * "_" * camelcase(name[idx:end])
  end
  return name
end
*/

import { options } from "./wailsjs/go/models";

type StringMap = {
  [key: string]: string;
};
type RGBObj = { r: number; g: number; b: number };

export default class Color {
  red: number;
  green: number;
  blue: number;
  opacity: number;
  hue: number;
  lightness: number;
  whiteness: number;
  sat: number;
  cyan: number;
  magenta: number;
  yellow: number;
  black: number;
  ncol: string;
  valid: boolean;

  static fromGO(col: options.RGBA): Color {
    return Color.rgba(col.r, col.g, col.b, col.a / 255);
  }
  toGO() {
    return new options.RGBA({
      r: this.red,
      g: this.green,
      b: this.blue,
      a: Math.round(this.opacity * 255),
    });
  }

  get blackness(): number {
    return 1 - this.lightness;
  }
  constructor(color: string | null = null) {
    this.red = 0;
    this.green = 0;
    this.blue = 0;
    this.hue = 0;
    this.lightness = 0;
    this.whiteness = 0;
    this.opacity = 1;
    this.sat = 0;
    this.cyan = 0;
    this.magenta = 0;
    this.yellow = 0;
    this.black = 0;
    this.ncol = "R";
    this.valid = false;
    if (color != null) {
      this.attachValues(Color.fromCSS(color));
    }
  }
  toRgbString() {
    return "rgb(" + this.red + ", " + this.green + ", " + this.blue + ")";
  }
  static rgba(r: number, g: number, b: number, a: number): Color {
    return new Color(`rgba(${r}, ${g}, ${b}, ${a}`);
  }
  toRgbaString() {
    return (
      "rgba(" +
      this.red +
      ", " +
      this.green +
      ", " +
      this.blue +
      ", " +
      this.opacity +
      ")"
    );
  }
  toHwbString() {
    return (
      "hwb(" +
      this.hue +
      ", " +
      Math.round(this.whiteness * 100) +
      "%, " +
      Math.round(this.blackness * 100) +
      "%)"
    );
  }
  toHwbStringDecimal() {
    return (
      "hwb(" + this.hue + ", " + this.whiteness + ", " + this.blackness + ")"
    );
  }
  toHwbaString() {
    return (
      "hwba(" +
      this.hue +
      ", " +
      Math.round(this.whiteness * 100) +
      "%, " +
      Math.round(this.blackness * 100) +
      "%, " +
      this.opacity +
      ")"
    );
  }
  toHslString() {
    return (
      "hsl(" +
      this.hue +
      ", " +
      Math.round(this.sat * 100) +
      "%, " +
      Math.round(this.lightness * 100) +
      "%)"
    );
  }
  toHslStringDecimal() {
    return "hsl(" + this.hue + ", " + this.sat + ", " + this.lightness + ")";
  }
  toHslaString() {
    return (
      "hsla(" +
      this.hue +
      ", " +
      Math.round(this.sat * 100) +
      "%, " +
      Math.round(this.lightness * 100) +
      "%, " +
      this.opacity +
      ")"
    );
  }
  toCmykString() {
    return (
      "cmyk(" +
      Math.round(this.cyan * 100) +
      "%, " +
      Math.round(this.magenta * 100) +
      "%, " +
      Math.round(this.yellow * 100) +
      "%, " +
      Math.round(this.black * 100) +
      "%)"
    );
  }
  toCmykStringDecimal() {
    return (
      "cmyk(" +
      this.cyan +
      ", " +
      this.magenta +
      ", " +
      this.yellow +
      ", " +
      this.black +
      ")"
    );
  }
  toNcolString() {
    return (
      this.ncol +
      ", " +
      Math.round(this.whiteness * 100) +
      "%, " +
      Math.round(this.blackness * 100) +
      "%"
    );
  }
  toNcolStringDecimal(): string {
    return this.ncol + ", " + this.whiteness + ", " + this.blackness;
  }
  toNcolaString(): string {
    return (
      this.ncol +
      ", " +
      Math.round(this.whiteness * 100) +
      "%, " +
      Math.round(this.blackness * 100) +
      "%, " +
      this.opacity
    );
  }
  toName(): string {
    var r: number, g: number, b: number, hex: string, name: string;
    for (name in Color.COLORS) {
      hex = Color.COLORS[name];
      r = parseInt(hex.substr(0, 2), 16);
      g = parseInt(hex.substr(2, 2), 16);
      b = parseInt(hex.substr(4, 2), 16);
      if (this.red == r && this.green == g && this.blue == b) {
        return name;
      }
    }
    return "";
  }
  toHexString() {
    var r = Color.toHex(this.red);
    var g = Color.toHex(this.green);
    var b = Color.toHex(this.blue);
    return "#" + r + g + b;
  }
  toRgb() {
    return { r: this.red, g: this.green, b: this.blue, a: this.opacity };
  }
  toHsl() {
    return { h: this.hue, s: this.sat, l: this.lightness, a: this.opacity };
  }
  toHwb() {
    return {
      h: this.hue,
      w: this.whiteness,
      b: this.blackness,
      a: this.opacity,
    };
  }
  toCmyk() {
    return {
      c: this.cyan,
      m: this.magenta,
      y: this.yellow,
      k: this.black,
      a: this.opacity,
    };
  }
  toNcol() {
    return {
      ncol: this.ncol,
      w: this.whiteness,
      b: this.blackness,
      a: this.opacity,
    };
  }
  isDark(n: number) {
    var m = n || 128;
    return (this.red * 299 + this.green * 587 + this.blue * 114) / 1000 < m;
  }
  lighter(n: number): Color {
    var x, rgb, color;
    x = n / 100 || 0.1;
    this.lightness += x;
    if (this.lightness > 1) {
      this.lightness = 1;
    }
    rgb = Color.hslToRgb(this.hue, this.sat, this.lightness);
    return Color.create(rgb, this.opacity, this.hue, this.sat);
  }
  darker(n: number): Color {
    var x, rgb;
    x = n / 100 || 0.1;
    this.lightness -= x;
    if (this.lightness < 0) {
      this.lightness = 0;
    }
    rgb = Color.hslToRgb(this.hue, this.sat, this.lightness);
    return Color.create(rgb, this.opacity, this.hue, this.sat);
  }
  attachValues(color: Color) {
    this.red = color.red;
    this.green = color.green;
    this.blue = color.blue;
    this.hue = color.hue;
    this.lightness = color.lightness;
    this.whiteness = color.whiteness;
    this.opacity = color.opacity;
    this.sat = color.sat;
    this.cyan = color.cyan;
    this.magenta = color.magenta;
    this.yellow = color.yellow;
    this.black = color.black;
    this.ncol = color.ncol;
    this.valid = color.valid;
  }
  static fromCSS(c: string): Color {
    var x: string,
      y: string,
      typ: string,
      arr: any[] = [],
      arrlength: number,
      i: number,
      opacity: boolean,
      match: any,
      a: number,
      hue: number = 0,
      sat: number = 0,
      rgb: RGBObj = { r: 0, g: 0, b: 0 };

    c = Color.trim(c.toLowerCase());
    x = c.substr(0, 1).toUpperCase();
    y = c.substr(1);
    a = 1;
    if (
      (x == "R" ||
        x == "Y" ||
        x == "G" ||
        x == "C" ||
        x == "B" ||
        x == "M" ||
        x == "W") &&
      !isNaN(Number(y))
    ) {
      c = "ncol(" + c + ")";
    }
    if (c.length != 3 && c.length != 6 && !isNaN(Number(c))) {
      c = "ncol(" + c + ")";
    }
    if (c.indexOf(",") > 0 && c.indexOf("(") == -1) {
      c = "ncol(" + c + ")";
    }
    if (
      c.substr(0, 3) == "rgb" ||
      c.substr(0, 3) == "hsl" ||
      c.substr(0, 3) == "hwb" ||
      c.substr(0, 4) == "ncol" ||
      c.substr(0, 4) == "cmyk"
    ) {
      if (c.substr(0, 4) == "ncol") {
        if (c.split(",").length == 4 && c.indexOf("ncola") == -1) {
          c = c.replace("ncol", "ncola");
        }
        typ = "ncol";
        c = c.substr(4);
      } else if (c.substr(0, 4) == "cmyk") {
        typ = "cmyk";
        c = c.substr(4);
      } else {
        typ = c.substr(0, 3);
        c = c.substr(3);
      }
      arrlength = 3;
      opacity = false;
      if (c.substr(0, 1).toLowerCase() == "a") {
        arrlength = 4;
        opacity = true;
        c = c.substr(1);
      } else if (typ == "cmyk") {
        arrlength = 4;
        if (c.split(",").length == 5) {
          arrlength = 5;
          opacity = true;
        }
      }
      c = c.replace("(", "");
      c = c.replace(")", "");
      arr = c.split(",");
      if (typ == "rgb") {
        if (arr.length != arrlength) {
          return Color.empty();
        }
        for (i = 0; i < arrlength; i++) {
          if (arr[i] == "" || arr[i] == " ") {
            arr[i] = "0";
          }
          if (arr[i].indexOf("%") > -1) {
            arr[i] = arr[i].replace("%", "");
            arr[i] = Number(parseInt(arr[i]) / 100).toFixed(2);
            if (i < 3) {
              arr[i] = Math.round(parseInt(arr[i]) * 255).toFixed(2);
            }
          }
          if (isNaN(Number(arr[i]))) {
            return Color.empty();
          }
          if (parseInt(arr[i]) > 255) {
            arr[i] = "255";
          }
          if (i < 3) {
            arr[i] = parseInt(arr[i]).toFixed(2);
          }
          if (i == 3 && Number(arr[i]) > 1) {
            arr[i] = "1";
          }
          rgb = { r: Number(arr[0]), g: Number(arr[1]), b: Number(arr[2]) };
          if (opacity == true) {
            a = Number(arr[3]);
          }
        }
      }
      if (typ == "hsl" || typ == "hwb" || typ == "ncol") {
        while (arr.length < arrlength) {
          arr.push("0");
        }
        if (typ == "hsl" || typ == "hwb") {
          if (parseInt(arr[0]) >= 360) {
            arr[0] = "0";
          }
        }
        for (i = 1; i < arrlength; i++) {
          if (arr[i].indexOf("%") > -1) {
            arr[i] = arr[i].replace("%", "");
            arr[i] = Number(arr[i]).toFixed(2);
            if (isNaN(Number(arr[i]))) {
              return Color.empty();
            }
            arr[i] = (parseInt(arr[i]) / 100).toFixed(2);
          }
          if (Number(arr[i]) > 1) {
            arr[i] = "1";
          }
          if (Number(arr[i]) < 0) {
            arr[i] = "0";
          }
        }
        if (typ == "hsl") {
          rgb = Color.hslToRgb(Number(arr[0]), Number(arr[1]), Number(arr[2]));
          hue = Number(arr[0]);
          sat = Number(arr[1]);
        } else if (typ == "hwb") {
          rgb = Color.hwbToRgb(Number(arr[0]), Number(arr[1]), Number(arr[2]));
        } else if (typ == "ncol") {
          const val = Color.ncolToRgb(arr[0], Number(arr[1]), Number(arr[2]));
          if (val != false) {
            rgb = val;
          }
        }
        if (opacity == true) {
          a = Number(arr[3]);
        }
      }
      if (typ == "cmyk") {
        while (arr.length < arrlength) {
          arr.push("0");
        }
        for (i = 0; i < arrlength; i++) {
          if (arr[i].indexOf("%") > -1) {
            arr[i] = arr[i].replace("%", "");
            if (isNaN(Number(arr[i]))) {
              return Color.empty();
            }
            arr[i] = (parseInt(arr[i]) / 100).toFixed(2);
          }
          if (Number(arr[i]) > 1) {
            arr[i] = "1";
          }
          if (Number(arr[i]) < 0) {
            arr[i] = "0";
          }
        }
        rgb = Color.cmykToRgb(
          Number(arr[0]),
          Number(arr[1]),
          Number(arr[2]),
          Number(arr[3]),
        );
        if (opacity == true) {
          a = Number(arr[4]);
        }
      }
    } else if (c.substr(0, 3) == "ncs") {
      const val = Color.ncsToRgb(c);
      if (val != false) {
        rgb = val;
      }
    } else {
      match = false;
      for (let colorname in Color.COLORS) {
        if (c.toLowerCase() == colorname.toLowerCase()) {
          var colorhex: string = Color.COLORS[colorname];
          match = true;
          rgb = {
            r: parseInt(colorhex.substr(0, 2), 16),
            g: parseInt(colorhex.substr(2, 2), 16),
            b: parseInt(colorhex.substr(4, 2), 16),
          };
          break;
        }
      }

      if (match == false) {
        c = c.replace("#", "");
        if (c.length == 3) {
          c =
            c.substr(0, 1) +
            c.substr(0, 1) +
            c.substr(1, 1) +
            c.substr(1, 1) +
            c.substr(2, 1) +
            c.substr(2, 1);
        }
        arr[0] = parseInt(c.substr(0, 2), 16).toFixed(2);
        arr[1] = parseInt(c.substr(2, 2), 16).toFixed(2);
        arr[2] = parseInt(c.substr(4, 2), 16).toFixed(2);
        for (i = 0; i < 3; i++) {
          if (isNaN(Number(arr[i]))) {
            return Color.empty();
          }
        }
        rgb = {
          r: Number(arr[0]),
          g: Number(arr[1]),
          b: Number(arr[2]),
        };
      }
    }
    return Color.create(rgb, a, hue, sat);
  }
  static create(
    rgb: { r: number; g: number; b: number },
    a: number,
    h: number,
    s: number,
  ): Color {
    var hsl, hwb, cmyk, ncol, color, hue, sat;
    if (!rgb) {
      return Color.empty();
    }
    if (!a) {
      a = 1;
    }
    hsl = Color.rgbToHsl(rgb.r, rgb.g, rgb.b);
    hwb = Color.rgbToHwb(rgb.r, rgb.g, rgb.b);
    cmyk = Color.rgbToCmyk(rgb.r, rgb.g, rgb.b);
    hue = h || hsl.h;
    sat = s || hsl.s;
    ncol = Color.hueToNcol(hue);
    color = new Color();
    color.red = rgb.r;
    color.green = rgb.g;
    color.blue = rgb.b;
    color.hue = hue;
    color.sat = sat;
    color.lightness = hsl.l;
    color.whiteness = hwb.w;
    color.cyan = cmyk.c;
    color.magenta = cmyk.m;
    color.yellow = cmyk.y;
    color.black = cmyk.k;
    if (ncol != undefined) color.ncol = ncol;
    color.opacity = a;
    color.valid = true;
    color.roundDecimals();
    return color;
  }
  static empty(): Color {
    return new Color();
  }

  roundDecimals() {
    this.red = Number(this.red.toFixed(0));
    this.green = Number(this.green.toFixed(0));
    this.blue = Number(this.blue.toFixed(0));
    this.hue = Number(this.hue.toFixed(0));
    this.sat = Number(this.sat.toFixed(2));
    this.lightness = Number(this.lightness.toFixed(2));
    this.whiteness = Number(this.whiteness.toFixed(2));
    this.cyan = Number(this.cyan.toFixed(2));
    this.magenta = Number(this.magenta.toFixed(2));
    this.yellow = Number(this.yellow.toFixed(2));
    this.black = Number(this.black.toFixed(2));
    this.ncol =
      this.ncol.substr(0, 1) + Math.round(Number(this.ncol.substr(1)));
    this.opacity = Number(this.opacity.toFixed(2));
  }
  static hslToRgb(
    hue: number,
    sat: number,
    light: number,
  ): { r: number; g: number; b: number } {
    var t1, t2, r, g, b;
    hue = hue / 60;
    if (light <= 0.5) {
      t2 = light * (sat + 1);
    } else {
      t2 = light + sat - light * sat;
    }
    t1 = light * 2 - t2;
    r = Color.hueToRgb(t1, t2, hue + 2) * 255;
    g = Color.hueToRgb(t1, t2, hue) * 255;
    b = Color.hueToRgb(t1, t2, hue - 2) * 255;
    return { r: r, g: g, b: b };
  }
  static hueToRgb(t1: number, t2: number, hue: number): number {
    // ?
    if (hue < 0) hue += 6;
    if (hue >= 6) hue -= 6;
    if (hue < 1) return (t2 - t1) * hue + t1;
    else if (hue < 3) return t2;
    else if (hue < 4) return (t2 - t1) * (4 - hue) + t1;
    else return t1;
  }
  static hwbToRgb(
    hue: number,
    white: number,
    black: number,
  ): { r: number; g: number; b: number } {
    var i: number,
      rgb: RGBObj,
      rgbArr: number[] = [];
    rgb = Color.hslToRgb(hue, 1, 0.5);
    rgbArr[0] = rgb.r / 255;
    rgbArr[1] = rgb.g / 255;
    rgbArr[2] = rgb.b / 255;
    for (i = 0; i < 3; i++) {
      rgbArr[i] *= 1 - white - black;
      rgbArr[i] += white;
      rgbArr[i] = Number(rgbArr[i] * 255);
    }
    return { r: rgbArr[0], g: rgbArr[1], b: rgbArr[2] };
  }
  static cmykToRgb(
    c: number,
    m: number,
    y: number,
    k: number,
  ): { r: number; g: number; b: number } {
    var r: number, g: number, b: number;
    r = 255 - Math.min(1, c * (1 - k) + k) * 255;
    g = 255 - Math.min(1, m * (1 - k) + k) * 255;
    b = 255 - Math.min(1, y * (1 - k) + k) * 255;
    return { r: r, g: g, b: b };
  }
  static ncolToRgb(ncol: string, white: number, black: number): RGBObj | false {
    var letter: string, percent: number, h: number, w: number, b: number;
    h = Number(ncol);
    if (isNaN(Number(ncol.substr(0, 1)))) {
      letter = ncol.substr(0, 1).toUpperCase();
      percent = Number(ncol.substr(1));
      if (isNaN(percent)) percent = 0;
      percent = Number(percent);
      if (isNaN(percent)) {
        return false;
      }
      if (letter == "R") {
        h = 0 + percent * 0.6;
      }
      if (letter == "Y") {
        h = 60 + percent * 0.6;
      }
      if (letter == "G") {
        h = 120 + percent * 0.6;
      }
      if (letter == "C") {
        h = 180 + percent * 0.6;
      }
      if (letter == "B") {
        h = 240 + percent * 0.6;
      }
      if (letter == "M") {
        h = 300 + percent * 0.6;
      }
      if (letter == "W") {
        h = 0;
        white = 1 - percent / 100;
        black = percent / 100;
      }
    }
    return Color.hwbToRgb(parseInt(h.toString()), white, black);
  }
  static hueToNcol(hue: number) {
    while (hue >= 360) {
      hue = hue - 360;
    }
    if (hue < 60) {
      return "R" + hue / 0.6;
    }
    if (hue < 120) {
      return "Y" + (hue - 60) / 0.6;
    }
    if (hue < 180) {
      return "G" + (hue - 120) / 0.6;
    }
    if (hue < 240) {
      return "C" + (hue - 180) / 0.6;
    }
    if (hue < 300) {
      return "B" + (hue - 240) / 0.6;
    }
    if (hue < 360) {
      return "M" + (hue - 300) / 0.6;
    }
  }
  static ncsToRgb(ncsstring: string) {
    var black,
      chroma,
      bc,
      percent,
      black1,
      chroma1,
      red1,
      factor1,
      blue1,
      red1,
      red2,
      green2,
      blue2,
      max,
      factor2,
      grey,
      r,
      g,
      b;
    var green1 = 0;
    ((red2 = 0),
      (blue2 = 0),
      (factor1 = 0),
      (factor2 = 0),
      (red1 = 0),
      (blue1 = 0));
    ncsstring = Color.trim(ncsstring).toUpperCase();
    ncsstring = ncsstring.replace("(", "");
    ncsstring = ncsstring.replace(")", "");
    ncsstring = ncsstring.replace("NCS", "NCS ");
    ncsstring = ncsstring.replace(/  /g, " ");
    if (ncsstring.indexOf("NCS") == -1) {
      ncsstring = "NCS " + ncsstring;
    }
    var ncs: RegExpMatchArray | null = ncsstring.match(
      /^(?:NCS|NCS\sS)\s(\d{2})(\d{2})-(N|[A-Z])(\d{2})?([A-Z])?$/,
    );
    if (ncs === null) return false;
    black = parseInt(ncs[1], 10);
    chroma = parseInt(ncs[2], 10);
    bc = ncs[3];
    if (bc != "N" && bc != "Y" && bc != "R" && bc != "B" && bc != "G") {
      return false;
    }
    percent = parseInt(ncs[4], 10) || 0;
    if (bc !== "N") {
      black1 = 1.05 * black - 5.25;
      chroma1 = chroma;
      if (bc === "Y" && percent <= 60) {
        red1 = 1;
      } else if (
        (bc === "Y" && percent > 60) ||
        (bc === "R" && percent <= 80)
      ) {
        if (bc === "Y") {
          factor1 = percent - 60;
        } else {
          factor1 = percent + 40;
        }
        red1 = (Math.sqrt(14884 - Math.pow(factor1, 2)) - 22) / 100;
      } else if ((bc === "R" && percent > 80) || bc === "B") {
        red1 = 0;
      } else if (bc === "G") {
        factor1 = percent - 170;
        red1 = (Math.sqrt(33800 - Math.pow(factor1, 2)) - 70) / 100;
      }
      if (bc === "Y" && percent <= 80) {
        blue1 = 0;
      } else if (
        (bc === "Y" && percent > 80) ||
        (bc === "R" && percent <= 60)
      ) {
        if (bc === "Y") {
          factor1 = percent - 80 + 20.5;
        } else {
          factor1 = percent + 20 + 20.5;
        }
        blue1 = (104 - Math.sqrt(11236 - Math.pow(factor1, 2))) / 100;
      } else if (
        (bc === "R" && percent > 60) ||
        (bc === "B" && percent <= 80)
      ) {
        if (bc === "R") {
          factor1 = percent - 60 - 60;
        } else {
          factor1 = percent + 40 - 60;
        }
        blue1 = (Math.sqrt(10000 - Math.pow(factor1, 2)) - 10) / 100;
      } else if (
        (bc === "B" && percent > 80) ||
        (bc === "G" && percent <= 40)
      ) {
        if (bc === "B") {
          factor1 = percent - 80 - 131;
        } else {
          factor1 = percent + 20 - 131;
        }
        blue1 = (122 - Math.sqrt(19881 - Math.pow(factor1, 2))) / 100;
      } else if (bc === "G" && percent > 40) {
        blue1 = 0;
      }
      if (bc === "Y") {
        green1 = (85 - (17 / 20) * percent) / 100;
      } else if (bc === "R" && percent <= 60) {
        green1 = 0;
      } else if (bc === "R" && percent > 60) {
        factor1 = percent - 60 + 35;
        green1 = (67.5 - Math.sqrt(5776 - Math.pow(factor1, 2))) / 100;
      } else if (bc === "B" && percent <= 60) {
        factor1 = 1 * percent - 68.5;
        green1 = (6.5 + Math.sqrt(7044.5 - Math.pow(factor1, 2))) / 100;
      } else if (
        (bc === "B" && percent > 60) ||
        (bc === "G" && percent <= 60)
      ) {
        green1 = 0.9;
      } else if (bc === "G" && percent > 60) {
        factor1 = percent - 60;
        green1 = (90 - (1 / 8) * factor1) / 100;
      }
      factor1 = (red1 + green1 + blue1) / 3;
      red2 = ((factor1 - red1) * (100 - chroma1)) / 100 + red1;
      green2 = ((factor1 - green1) * (100 - chroma1)) / 100 + green1;
      blue2 = ((factor1 - blue1) * (100 - chroma1)) / 100 + blue1;
      if (red2 > green2 && red2 > blue2) {
        max = red2;
      } else if (green2 > red2 && green2 > blue2) {
        max = green2;
      } else if (blue2 > red2 && blue2 > green2) {
        max = blue2;
      } else {
        max = (red2 + green2 + blue2) / 3;
      }
      factor2 = 1 / max;
      r = ((red2 * factor2 * (100 - black1)) / 100) * 255;
      g = ((green2 * factor2 * (100 - black1)) / 100) * 255;
      b = ((blue2 * factor2 * (100 - black1)) / 100) * 255;
      if (r > 255) {
        r = 255;
      }
      if (g > 255) {
        g = 255;
      }
      if (b > 255) {
        b = 255;
      }
      if (r < 0) {
        r = 0;
      }
      if (g < 0) {
        g = 0;
      }
      if (b < 0) {
        b = 0;
      }
    } else {
      grey = (1 - black / 100) * 255;
      if (grey > 255) {
        grey = 255;
      }
      if (grey < 0) {
        grey = 0;
      }
      r = grey;
      g = grey;
      b = grey;
    }
    return {
      r: r,
      g: g,
      b: b,
    };
  }
  static rgbToHsl(
    r: number,
    g: number,
    b: number,
  ): { h: number; s: number; l: number } {
    var min: number,
      max: number,
      i: number,
      l: number,
      s: number,
      maxcolor: number,
      h: number = 0,
      rgb: number[] = [];
    rgb[0] = r / 255;
    rgb[1] = g / 255;
    rgb[2] = b / 255;
    min = rgb[0];
    max = rgb[0];
    maxcolor = 0;
    for (i = 0; i < rgb.length - 1; i++) {
      if (rgb[i + 1] <= min) {
        min = rgb[i + 1];
      }
      if (rgb[i + 1] >= max) {
        max = rgb[i + 1];
        maxcolor = i + 1;
      }
    }
    if (maxcolor == 0) {
      h = (rgb[1] - rgb[2]) / (max - min);
    } else if (maxcolor == 1) {
      h = 2 + (rgb[2] - rgb[0]) / (max - min);
    } else if (maxcolor == 2) {
      h = 4 + (rgb[0] - rgb[1]) / (max - min);
    }
    if (isNaN(Number(h))) {
      h = 0;
    }
    h = h * 60;
    if (h < 0) {
      h = h + 360;
    }
    l = (min + max) / 2;
    if (min == max) {
      s = 0;
    } else {
      if (l < 0.5) {
        s = (max - min) / (max + min);
      } else {
        s = (max - min) / (2 - max - min);
      }
    }
    s = s;
    return { h: h, s: s, l: l };
  }
  static rgbToHwb(
    r: number,
    g: number,
    b: number,
  ): { h: number; w: number; b: number } {
    var h: number, w: number, b: number;
    r = r / 255;
    g = g / 255;
    b = b / 255;
    var max: number = Math.max(r, g, b),
      min: number = Math.min(r, g, b);
    var chroma: number = max - min;
    if (chroma == 0) {
      h = 0;
    } else if (r == max) {
      h = (((g - b) / chroma) % 6) * 360;
    } else if (g == max) {
      h = (((b - r) / chroma + 2) % 6) * 360;
    } else {
      h = (((r - g) / chroma + 4) % 6) * 360;
    }
    w = min;
    b = 1 - max;
    return { h: h, w: w, b: b };
  }
  static rgbToCmyk(
    r: number,
    g: number,
    b: number,
  ): { c: number; m: number; y: number; k: number } {
    var c: number, m: number, y: number, k: number;
    r = r / 255;
    g = g / 255;
    b = b / 255;
    var max: number = Math.max(r, g, b);
    k = 1 - max;
    if (k == 1) {
      c = 0;
      m = 0;
      y = 0;
    } else {
      c = (1 - r - k) / (1 - k);
      m = (1 - g - k) / (1 - k);
      y = (1 - b - k) / (1 - k);
    }
    return { c: c, m: m, y: y, k: k };
  }
  static toHex(n: number): string {
    var hex = n.toString(16);
    while (hex.length < 2) {
      hex = "0" + hex;
    }
    return hex;
  }
  static trim(x: string): string {
    return x.replace(/^\s+|\s+$/g, "");
  }
  public static readonly COLORS: StringMap = {
    AliceBlue: "f0f8ff",
    AntiqueWhite: "faebd7",
    Aqua: "00ffff",
    Aquamarine: "7fffd4",
    Azure: "f0ffff",
    Beige: "f5f5dc",
    Bisque: "ffe4c4",
    Black: "000000",
    BlanchedAlmond: "ffebcd",
    Blue: "0000ff",
    BlueViolet: "8a2be2",
    Brown: "a52a2a",
    BurlyWood: "deb887",
    CadetBlue: "5f9ea0",
    Chartreuse: "7fff00",
    Chocolate: "d2691e",
    Coral: "ff7f50",
    CornflowerBlue: "6495ed",
    Cornsilk: "fff8dc",
    Crimson: "dc143c",
    Cyan: "00ffff",
    DarkBlue: "00008b",
    DarkCyan: "008b8b",
    DarkGoldenRod: "b8860b",
    DarkGray: "a9a9a9",
    DarkGrey: "a9a9a9",
    DarkGreen: "006400",
    DarkKhaki: "bdb76b",
    DarkMagenta: "8b008b",
    DarkOliveGreen: "556b2f",
    DarkOrange: "ff8c00",
    DarkOrchid: "9932cc",
    DarkRed: "8b0000",
    DarkSalmon: "e9967a",
    DarkSeaGreen: "8fbc8f",
    DarkSlateBlue: "483d8b",
    DarkSlateGray: "2f4f4f",
    DarkSlateGrey: "2f4f4f",
    DarkTurquoise: "00ced1",
    DarkViolet: "9400d3",
    DeepPink: "ff1493",
    DeepSkyBlue: "00bfff",
    DimGray: "696969",
    DimGrey: "696969",
    DodgerBlue: "1e90ff",
    FireBrick: "b22222",
    FloralWhite: "fffaf0",
    ForestGreen: "228b22",
    Fuchsia: "ff00ff",
    Gainsboro: "dcdcdc",
    GhostWhite: "f8f8ff",
    Gold: "ffd700",
    GoldenRod: "daa520",
    Gray: "808080",
    Grey: "808080",
    Green: "008000",
    GreenYellow: "adff2f",
    HoneyDew: "f0fff0",
    HotPink: "ff69b4",
    IndianRed: "cd5c5c",
    Indigo: "4b0082",
    Ivory: "fffff0",
    Khaki: "f0e68c",
    Lavender: "e6e6fa",
    LavenderBlush: "fff0f5",
    LawnGreen: "7cfc00",
    LemonChiffon: "fffacd",
    LightBlue: "add8e6",
    LightCoral: "f08080",
    LightCyan: "e0ffff",
    LightGoldenRodYellow: "fafad2",
    LightGray: "d3d3d3",
    LightGrey: "d3d3d3",
    LightGreen: "90ee90",
    LightPink: "ffb6c1",
    LightSalmon: "ffa07a",
    LightSeaGreen: "20b2aa",
    LightSkyBlue: "87cefa",
    LightSlateGray: "778899",
    LightSlateGrey: "778899",
    LightSteelBlue: "b0c4de",
    LightYellow: "ffffe0",
    Lime: "00ff00",
    LimeGreen: "32cd32",
    Linen: "faf0e6",
    Magenta: "ff00ff",
    Maroon: "800000",
    MediumAquaMarine: "66cdaa",
    MediumBlue: "0000cd",
    MediumOrchid: "ba55d3",
    MediumPurple: "9370db",
    MediumSeaGreen: "3cb371",
    MediumSlateBlue: "7b68ee",
    MediumSpringGreen: "00fa9a",
    MediumTurquoise: "48d1cc",
    MediumVioletRed: "c71585",
    MidnightBlue: "191970",
    MintCream: "f5fffa",
    MistyRose: "ffe4e1",
    Moccasin: "ffe4b5",
    NavajoWhite: "ffdead",
    Navy: "000080",
    OldLace: "fdf5e6",
    Olive: "808000",
    OliveDrab: "6b8e23",
    Orange: "ffa500",
    OrangeRed: "ff4500",
    Orchid: "da70d6",
    PaleGoldenRod: "eee8aa",
    PaleGreen: "98fb98",
    PaleTurquoise: "afeeee",
    PaleVioletRed: "db7093",
    PapayaWhip: "ffefd5",
    PeachPuff: "ffdab9",
    Peru: "cd853f",
    Pink: "ffc0cb",
    Plum: "dda0dd",
    PowderBlue: "b0e0e6",
    Purple: "800080",
    RebeccaPurple: "663399",
    Red: "ff0000",
    RosyBrown: "bc8f8f",
    RoyalBlue: "4169e1",
    SaddleBrown: "8b4513",
    Salmon: "fa8072",
    SandyBrown: "f4a460",
    SeaGreen: "2e8b57",
    SeaShell: "fff5ee",
    Sienna: "a0522d",
    Silver: "c0c0c0",
    SkyBlue: "87ceeb",
    SlateBlue: "6a5acd",
    SlateGray: "708090",
    SlateGrey: "708090",
    Snow: "fffafa",
    SpringGreen: "00ff7f",
    SteelBlue: "4682b4",
    Tan: "d2b48c",
    Teal: "008080",
    Thistle: "d8bfd8",
    Tomato: "ff6347",
    Turquoise: "40e0d0",
    Violet: "ee82ee",
    Wheat: "f5deb3",
    White: "ffffff",
    WhiteSmoke: "f5f5f5",
    Yellow: "ffff00",
    YellowGreen: "9acd32",
  };
}
