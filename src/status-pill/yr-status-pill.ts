// <yr-status-pill status="active">Running</yr-status-pill> — a live status
// indicator using the shared signal-* tokens (active/pending/stopped/idle,
// the same railway-signal vocabulary tokens.css already defines). Light DOM,
// styled by yr-status-pill.css; purely presentational, no behavior yet.
export class YrStatusPill extends HTMLElement {
  static observedAttributes = ['status'];
}

customElements.define('yr-status-pill', YrStatusPill);
