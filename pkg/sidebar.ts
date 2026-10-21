export class YrSidebar extends HTMLElement {
  static observedAttributes = ['collapsed'];

  attributeChangedCallback() {
    // CSS handles all visual changes via yr-sidebar[collapsed]
  }
}

// SidebarHeader — toggle button is server-rendered by templ. This class
// only wires up the click handler.
export class YrSidebarHeader extends HTMLElement {
  connectedCallback() {
    const toggle = this.querySelector('.yr-sidebar-toggle');
    if (toggle) {
      toggle.addEventListener('click', () => {
        const sidebar = this.closest('yr-sidebar');
        if (sidebar) sidebar.toggleAttribute('collapsed');
      });
    }
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

// SidebarOrg — structure is server-rendered by templ.
export class YrSidebarOrg extends HTMLElement {}

customElements.define('yr-sidebar', YrSidebar);
customElements.define('yr-sidebar-header', YrSidebarHeader);
customElements.define('yr-sidebar-nav', YrSidebarNav);
customElements.define('yr-sidebar-footer', YrSidebarFooter);
customElements.define('yr-sidebar-link', YrSidebarLink);
customElements.define('yr-sidebar-org', YrSidebarOrg);
