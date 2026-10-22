export class YrCappyDrawer extends HTMLElement {
  static observedAttributes = ['open'];

  private onKeyDown = (e: KeyboardEvent) => {
    if (e.key === 'Escape' && this.hasAttribute('open')) {
      this.removeAttribute('open');
    }
  };

  connectedCallback() {
    document.addEventListener('keydown', this.onKeyDown);
    this.querySelector('.yr-cappy-drawer__close')?.addEventListener('click', this.close);
    this.querySelector('.yr-cappy-drawer__collapse')?.addEventListener('click', this.toggleCollapse);
    this.querySelector('.yr-cappy-drawer__escalate')?.addEventListener('click', this.escalate);
    this.querySelector('.yr-cappy-drawer__log')?.addEventListener('click', this.onLogClick);
  }

  disconnectedCallback() {
    document.removeEventListener('keydown', this.onKeyDown);
    this.querySelector('.yr-cappy-drawer__close')?.removeEventListener('click', this.close);
    this.querySelector('.yr-cappy-drawer__collapse')?.removeEventListener('click', this.toggleCollapse);
    this.querySelector('.yr-cappy-drawer__escalate')?.removeEventListener('click', this.escalate);
    this.querySelector('.yr-cappy-drawer__log')?.removeEventListener('click', this.onLogClick);
  }

  attributeChangedCallback() {
    // CSS handles visual changes via yr-cappy-drawer[open]
  }

  private close = () => {
    this.removeAttribute('open');
  };

  private toggleCollapse = () => {
    this.toggleAttribute('collapsed');
  };

  private escalate = () => {
    this.dispatchEvent(new CustomEvent('yr-chat-escalate', { bubbles: true }));
  };

  private onLogClick = (e: Event) => {
    const msg = (e.target as HTMLElement).closest('.yr-chat-msg');
    if (!msg) return;
    msg.toggleAttribute('data-show-time');
  };
}

customElements.define('yr-cappy-drawer', YrCappyDrawer);
