// <yr-button variant="primary" type="submit">Save</yr-button> — light DOM, no
// shadow root. Styled globally by yr-button.css targeting the tag/attribute
// directly, so it renders correctly the instant the server's HTML arrives.
// A custom element has none of a native <button>'s built-in behavior, so this
// adds the baseline: button role, keyboard activation, and form submission
// when type="submit" inside a <form> (e.g. depot's "Add to cart").
export class YrButton extends HTMLElement {
  static observedAttributes = ['disabled'];

  connectedCallback() {
    if (!this.hasAttribute('role')) this.setAttribute('role', 'button');
    this.setAttribute('tabindex', this.disabled ? '-1' : '0');
    this.addEventListener('click', this.handleActivate);
    this.addEventListener('keydown', this.handleKeydown);
  }

  disconnectedCallback() {
    this.removeEventListener('click', this.handleActivate);
    this.removeEventListener('keydown', this.handleKeydown);
  }

  attributeChangedCallback(name: string) {
    if (name !== 'disabled') return;
    this.setAttribute('aria-disabled', String(this.disabled));
    this.setAttribute('tabindex', this.disabled ? '-1' : '0');
  }

  get disabled(): boolean {
    return this.hasAttribute('disabled');
  }

  private handleActivate = (event: Event) => {
    if (this.disabled) {
      event.preventDefault();
      return;
    }
    if (this.getAttribute('type') === 'submit') {
      this.closest('form')?.requestSubmit();
    }
  };

  private handleKeydown = (event: KeyboardEvent) => {
    if ((event.key === 'Enter' || event.key === ' ') && !this.disabled) {
      event.preventDefault();
      this.click();
    }
  };
}

customElements.define('yr-button', YrButton);
