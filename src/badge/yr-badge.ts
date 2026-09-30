// <yr-badge tone="navy">Official</yr-badge> — a static label pill. The tone
// is a generic palette slot (navy/amber/blue/pink/green/purple/slate); the
// caller maps its own semantics (provenance, category, kind, ...) onto a
// tone rather than the component knowing about "hr" or "official". Light
// DOM, styled by yr-badge.css; purely presentational, no behavior.
export class YrBadge extends HTMLElement {
  static observedAttributes = ['tone'];
}

customElements.define('yr-badge', YrBadge);
