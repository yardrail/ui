// <yr-alert tone="success"><svg>...</svg><strong>Added!</strong></yr-alert> —
// a banner/flash message. Content (icon + text) is light DOM supplied by the
// caller. Light DOM, styled by yr-alert.css; purely presentational, no
// behavior.
export class YrAlert extends HTMLElement {
  static observedAttributes = ['tone'];
}

customElements.define('yr-alert', YrAlert);
