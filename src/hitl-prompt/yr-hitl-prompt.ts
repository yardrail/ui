// <yr-hitl-prompt tone="warning"><div class="yr-hitl-prompt-title">Approve?
// </div>...</yr-hitl-prompt> — a human-in-the-loop prompt card. Shows a
// question, context, and response controls. Sets role="alert" automatically
// for screen reader urgency. Light DOM, styled by yr-hitl-prompt.css.
export class YrHitlPrompt extends HTMLElement {
  static observedAttributes = ['tone'];

  connectedCallback() {
    if (!this.getAttribute('role')) this.setAttribute('role', 'alert');
  }
}

customElements.define('yr-hitl-prompt', YrHitlPrompt);
