// <yr-tabs><button slot="tab" data-tab="overview">Overview</button>...
// <div slot="panel" data-tab="overview">...</div>...</yr-tabs> — an ARIA
// tabs widget (WAI-ARIA APG "tabs, automatic activation" pattern). Tab
// buttons and panels are light-DOM, service-authored, and paired up by a
// shared `data-tab` key; the component only wires up the tablist wrapper,
// roles/aria-*, keyboard nav, and — when the `sync-hash` attribute is
// present — the active tab is mirrored to `location.hash` so it survives
// refresh/back-forward and is directly linkable.
let instanceCount = 0;

export class YrTabs extends HTMLElement {
  private tabs: HTMLElement[] = [];
  private panels: HTMLElement[] = [];

  connectedCallback() {
    this.tabs = Array.from(this.querySelectorAll(':scope > [slot="tab"]'));
    this.panels = Array.from(this.querySelectorAll(':scope > [slot="panel"]'));
    if (!this.tabs.length) return;

    if (!this.id) this.id = `yr-tabs-${++instanceCount}`;

    const tablist = document.createElement('div');
    tablist.setAttribute('role', 'tablist');
    const label = this.getAttribute('label');
    if (label) tablist.setAttribute('aria-label', label);
    this.tabs.forEach((tab) => tablist.appendChild(tab));
    this.prepend(tablist);

    this.tabs.forEach((tab, i) => {
      const key = tab.dataset.tab ?? String(i);
      tab.id = `${this.id}-tab-${key}`;
      tab.setAttribute('role', 'tab');
      tab.setAttribute('aria-controls', `${this.id}-panel-${key}`);
      tab.addEventListener('click', () => this.select(key));
      tab.addEventListener('keydown', this.handleKeydown);

      const panel = this.panels.find((p) => p.dataset.tab === key);
      if (panel) {
        panel.id = `${this.id}-panel-${key}`;
        panel.setAttribute('role', 'tabpanel');
        panel.setAttribute('aria-labelledby', tab.id);
        panel.tabIndex = 0;
      }
    });

    if (this.syncHash) window.addEventListener('hashchange', this.handleHashChange);

    const hashKey = location.hash.slice(1);
    const initial = this.syncHash && this.tabs.some((t) => t.dataset.tab === hashKey)
      ? hashKey
      : this.tabs[0].dataset.tab ?? '';
    this.render(initial);
  }

  disconnectedCallback() {
    if (this.syncHash) window.removeEventListener('hashchange', this.handleHashChange);
  }

  private get syncHash(): boolean {
    return this.hasAttribute('sync-hash');
  }

  private select(key: string) {
    if (this.syncHash) {
      // Triggers the hashchange listener below, which does the actual render.
      location.hash = key;
    } else {
      this.render(key);
    }
  }

  private render(key: string) {
    this.tabs.forEach((tab) => {
      const selected = tab.dataset.tab === key;
      tab.setAttribute('aria-selected', String(selected));
      tab.tabIndex = selected ? 0 : -1;
    });
    this.panels.forEach((panel) => {
      panel.hidden = panel.dataset.tab !== key;
    });
  }

  private handleHashChange = () => {
    const key = location.hash.slice(1);
    if (this.tabs.some((t) => t.dataset.tab === key)) this.render(key);
  };

  private handleKeydown = (event: KeyboardEvent) => {
    const currentIndex = this.tabs.findIndex((t) => t.id === (event.target as HTMLElement).id);
    if (currentIndex === -1) return;

    let nextIndex: number | null = null;
    switch (event.key) {
      case 'ArrowRight':
        nextIndex = (currentIndex + 1) % this.tabs.length;
        break;
      case 'ArrowLeft':
        nextIndex = (currentIndex - 1 + this.tabs.length) % this.tabs.length;
        break;
      case 'Home':
        nextIndex = 0;
        break;
      case 'End':
        nextIndex = this.tabs.length - 1;
        break;
      default:
        return;
    }

    event.preventDefault();
    const next = this.tabs[nextIndex];
    this.select(next.dataset.tab ?? '');
    next.focus();
  };
}

customElements.define('yr-tabs', YrTabs);
