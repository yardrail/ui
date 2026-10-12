// <yr-status-banner tone="success"> — a persistent status display bar.
// Similar to yr-alert but for workflow status rather than notifications.
// Tone sets the color; light-DOM children include a pill, message, and
// action buttons. Purely presentational.
export class YrStatusBanner extends HTMLElement {
  static observedAttributes = ['tone'];
}

customElements.define('yr-status-banner', YrStatusBanner);
