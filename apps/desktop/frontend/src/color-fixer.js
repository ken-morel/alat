function oklchToRgb(l, c, h) {
  l /= 100;
  c /= 100;
  h /= 360;

  const l_ = l;
  const a_ = c * Math.cos(2 * Math.PI * h);
  const b_ = c * Math.sin(2 * Math.PI * h);

  const l_2 = Math.pow(l_ + 0.3963377774 * a_ + 0.2158037573 * b_, 3);
  const m_2 = Math.pow(l_ - 0.1055613458 * a_ - 0.0638541728 * b_, 3);
  const s_2 = Math.pow(l_ - 0.0894841775 * a_ - 1.291485548 * b_, 3);

  let r = 4.0767416621 * l_2 - 3.3077115913 * m_2 + 0.2309699292 * s_2;
  let g = -1.2684380046 * l_2 + 2.6097574011 * m_2 - 0.3413193965 * s_2;
  let b = -0.0041960863 * l_2 - 0.7034186147 * m_2 + 1.707614701 * s_2;

  r = Math.max(0, Math.min(1, r));
  g = Math.max(0, Math.min(1, g));
  b = Math.max(0, Math.min(1, b));

  return [Math.round(r * 255), Math.round(g * 255), Math.round(b * 255)];
}

function convertOklchInCss() {
  const oklchRegex = /oklch\(([^)]+)\)/;
  const sheets = Array.from(document.styleSheets);

  for (const sheet of sheets) {
    try {
      const rules = Array.from(sheet.cssRules);
      for (const rule of rules) {
        if (rule.style) {
          for (const prop of Array.from(rule.style)) {
            const value = rule.style.getPropertyValue(prop);
            const match = value.match(oklchRegex);
            if (match) {
              const parts = match[1].split(" ");
              const l = parseFloat(parts[0]);
              const c = parseFloat(parts[1]);
              const h = parseFloat(parts[2]);
              const [r, g, b] = oklchToRgb(l, c, h);
              const rgbaValue = `rgba(${r}, ${g}, ${b}, 1)`;
              rule.style.setProperty(prop, rgbaValue);
            }
          }
        }
      }
    } catch (e) {
      // Ignore CORS errors on external stylesheets
    }
  }
}

// setInterval(convertOklchInCss, 1000);

