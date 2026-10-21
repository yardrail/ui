export class YrSidebar extends HTMLElement {
  static observedAttributes = ['collapsed'];

  attributeChangedCallback() {
    // CSS handles all visual changes via yr-sidebar[collapsed]
  }
}

export class YrSidebarHeader extends HTMLElement {
  connectedCallback() {
    const toggle = document.createElement('button');
    toggle.className = 'yr-sidebar-toggle';
    toggle.setAttribute('aria-label', 'Toggle sidebar');
    toggle.innerHTML =
      `<svg class="yr-sidebar-toggle-close" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M9 3v18"/><path d="m16 15-3-3 3-3"/></svg>` +
      `<svg class="yr-sidebar-toggle-open" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M9 3v18"/><path d="m14 9 3 3-3 3"/></svg>`;
    toggle.addEventListener('click', () => {
      const sidebar = this.closest('yr-sidebar');
      if (sidebar) sidebar.toggleAttribute('collapsed');
    });
    this.insertBefore(toggle, this.firstChild);
  }
}
export class YrSidebarNav extends HTMLElement {}
export class YrSidebarFooter extends HTMLElement {}

export class YrSidebarLink extends HTMLElement {
  static observedAttributes = ['href', 'active'];

  private onClick = (e: Event) => {
    const href = this.getAttribute('href');
    if (!href) return;
    e.preventDefault();
    window.location.href = href;
  };

  connectedCallback() {
    this.setAttribute('role', 'link');
    this.setAttribute('tabindex', '0');
    this.addEventListener('click', this.onClick);
    this.addEventListener('keydown', this.onKeyDown);
  }

  disconnectedCallback() {
    this.removeEventListener('click', this.onClick);
    this.removeEventListener('keydown', this.onKeyDown);
  }

  private onKeyDown = (e: KeyboardEvent) => {
    if (e.key === 'Enter') this.click();
  };

  attributeChangedCallback() {
    // CSS handles visual changes via yr-sidebar-link[active]
  }
}

export class YrSidebarOrg extends HTMLElement {
  static observedAttributes = ['name', 'initial'];

  connectedCallback() {
    this.render();
  }

  attributeChangedCallback() {
    this.render();
  }

  private render() {
    const name = this.getAttribute('name') || '';
    const initial = this.getAttribute('initial') || name.charAt(0).toUpperCase();
    this.innerHTML =
      `<span class="yr-sidebar-org-pfp">${initial}</span>` +
      `<span class="yr-sidebar-org-name">${name}</span>`;
  }
}

customElements.define('yr-sidebar', YrSidebar);
customElements.define('yr-sidebar-header', YrSidebarHeader);
customElements.define('yr-sidebar-nav', YrSidebarNav);
customElements.define('yr-sidebar-footer', YrSidebarFooter);
customElements.define('yr-sidebar-link', YrSidebarLink);
customElements.define('yr-sidebar-org', YrSidebarOrg);
