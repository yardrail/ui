// <yr-pill tone="navy">Official</yr-pill> — a colored label. Tone sets the
// color; shape="pill" makes it fully rounded. Text content is the label;
// children can include SVG icons. Light DOM, styled by yr-pill.css.
export class YrPill extends HTMLElement {
  static observedAttributes = ['tone', 'shape'];
}

customElements.define('yr-pill', YrPill);
