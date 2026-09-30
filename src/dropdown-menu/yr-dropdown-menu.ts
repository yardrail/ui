// <yr-dropdown-menu><button slot="trigger">...</button><div slot="panel">
// ...</div></yr-dropdown-menu> — a trigger + panel disclosure. Centralizes
// the outside-click-to-close JS that depot's mkt_layout.templ currently
// inlines by hand, and replaces yardrail's structurally-identical
// .profile-menu <details> disclosure. Light-DOM children (trigger + panel
// content) stay fully service-authored.
export class YrDropdownMenu extends HTMLElement {
  private trigger: Element | null = null;

  connectedCallback() {
    this.trigger = this.querySelector('[slot="trigger"]');
    this.trigger?.addEventListener('click', this.handleTriggerClick);
    document.addEventListener('click', this.handleDocumentClick);
    document.addEventListener('keydown', this.handleKeydown);
  }

  disconnectedCallback() {
    this.trigger?.removeEventListener('click', this.handleTriggerClick);
    document.removeEventListener('click', this.handleDocumentClick);
    document.removeEventListener('keydown', this.handleKeydown);
  }

  private handleTriggerClick = (event: Event) => {
    event.stopPropagation();
    this.toggleAttribute('open');
  };

  private handleDocumentClick = (event: Event) => {
    if (this.hasAttribute('open') && !this.contains(event.target as Node)) {
      this.removeAttribute('open');
    }
  };

  private handleKeydown = (event: KeyboardEvent) => {
    if (event.key === 'Escape') this.removeAttribute('open');
  };
}

customElements.define('yr-dropdown-menu', YrDropdownMenu);
