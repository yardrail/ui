// pkg/button.ts
var YrButton = class extends HTMLElement {
  static observedAttributes = ["disabled"];
  connectedCallback() {
    if (!this.hasAttribute("role")) this.setAttribute("role", "button");
    this.setAttribute("tabindex", this.disabled ? "-1" : "0");
    this.addEventListener("click", this.handleActivate);
    this.addEventListener("keydown", this.handleKeydown);
  }
  disconnectedCallback() {
    this.removeEventListener("click", this.handleActivate);
    this.removeEventListener("keydown", this.handleKeydown);
  }
  attributeChangedCallback(name) {
    if (name !== "disabled") return;
    this.setAttribute("aria-disabled", String(this.disabled));
    this.setAttribute("tabindex", this.disabled ? "-1" : "0");
  }
  get disabled() {
    return this.hasAttribute("disabled");
  }
  handleActivate = (event) => {
    if (this.disabled) {
      event.preventDefault();
      return;
    }
    if (this.getAttribute("type") === "submit") {
      this.closest("form")?.requestSubmit();
    }
  };
  handleKeydown = (event) => {
    if ((event.key === "Enter" || event.key === " ") && !this.disabled) {
      event.preventDefault();
      this.click();
    }
  };
};
customElements.define("yr-button", YrButton);

// pkg/pill.ts
var YrPill = class extends HTMLElement {
  static observedAttributes = ["tone", "shape"];
};
customElements.define("yr-pill", YrPill);

// pkg/avatar.ts
var YrAvatar = class extends HTMLElement {
  static observedAttributes = ["size"];
};
customElements.define("yr-avatar", YrAvatar);

// pkg/alert.ts
var YrAlert = class extends HTMLElement {
  static observedAttributes = ["tone"];
};
customElements.define("yr-alert", YrAlert);

// pkg/stat.ts
var YrStat = class extends HTMLElement {
};
customElements.define("yr-stat", YrStat);

// pkg/rating.ts
var YrRating = class extends HTMLElement {
};
customElements.define("yr-rating", YrRating);

// pkg/empty_state.ts
var YrEmptyState = class extends HTMLElement {
};
customElements.define("yr-empty-state", YrEmptyState);

// pkg/copy_field.ts
var YrCopyField = class extends HTMLElement {
  resetTimer;
  connectedCallback() {
    const button = this.querySelector(".yr-copy-field-button");
    if (button) {
      button.addEventListener("click", () => {
        const value = button.dataset.value ?? "";
        this.handleCopy(value, button);
      });
    }
  }
  disconnectedCallback() {
    if (this.resetTimer !== void 0) window.clearTimeout(this.resetTimer);
  }
  async handleCopy(value, button) {
    try {
      await navigator.clipboard.writeText(value);
      button.textContent = "Copied!";
    } catch {
      button.textContent = "Copy failed";
    }
    if (this.resetTimer !== void 0) window.clearTimeout(this.resetTimer);
    this.resetTimer = window.setTimeout(() => {
      button.textContent = "Copy";
    }, 1500);
  }
};
customElements.define("yr-copy-field", YrCopyField);

// pkg/dropdown_menu.ts
var YrDropdownMenu = class extends HTMLElement {
  trigger = null;
  connectedCallback() {
    this.trigger = this.querySelector('[slot="trigger"]');
    this.trigger?.addEventListener("click", this.handleTriggerClick);
    document.addEventListener("click", this.handleDocumentClick);
    document.addEventListener("keydown", this.handleKeydown);
  }
  disconnectedCallback() {
    this.trigger?.removeEventListener("click", this.handleTriggerClick);
    document.removeEventListener("click", this.handleDocumentClick);
    document.removeEventListener("keydown", this.handleKeydown);
  }
  handleTriggerClick = (event) => {
    event.stopPropagation();
    this.toggleAttribute("open");
  };
  handleDocumentClick = (event) => {
    if (this.hasAttribute("open") && !this.contains(event.target)) {
      this.removeAttribute("open");
    }
  };
  handleKeydown = (event) => {
    if (event.key === "Escape") this.removeAttribute("open");
  };
};
customElements.define("yr-dropdown-menu", YrDropdownMenu);

// pkg/account_menu.ts
var YrAccountMenu = class extends YrDropdownMenu {
};
customElements.define("yr-account-menu", YrAccountMenu);

// pkg/live_duration.ts
var YrLiveDuration = class extends HTMLElement {
  static observedAttributes = ["started-at"];
  intervalId;
  connectedCallback() {
    this.tick();
    this.intervalId = window.setInterval(() => this.tick(), 1e3);
  }
  disconnectedCallback() {
    if (this.intervalId !== void 0) window.clearInterval(this.intervalId);
  }
  attributeChangedCallback() {
    this.tick();
  }
  tick() {
    const startedAt = this.getAttribute("started-at");
    if (!startedAt) {
      this.textContent = "--:--:--";
      return;
    }
    const elapsedMs = Math.max(0, Date.now() - new Date(startedAt).getTime());
    this.textContent = formatDuration(elapsedMs);
  }
};
function formatDuration(ms) {
  const totalSeconds = Math.floor(ms / 1e3);
  const h = Math.floor(totalSeconds / 3600);
  const m = Math.floor(totalSeconds % 3600 / 60);
  const s = totalSeconds % 60;
  return [h, m, s].map((n) => String(n).padStart(2, "0")).join(":");
}
customElements.define("yr-live-duration", YrLiveDuration);

// pkg/tabs.ts
var instanceCount = 0;
var YrTabs = class extends HTMLElement {
  tabs = [];
  panels = [];
  connectedCallback() {
    this.tabs = Array.from(this.querySelectorAll(':scope > [slot="tab"]'));
    this.panels = Array.from(this.querySelectorAll(':scope > [slot="panel"]'));
    if (!this.tabs.length) return;
    if (!this.id) this.id = `yr-tabs-${++instanceCount}`;
    const tablist = document.createElement("div");
    tablist.setAttribute("role", "tablist");
    const label = this.getAttribute("label");
    if (label) tablist.setAttribute("aria-label", label);
    this.tabs.forEach((tab) => tablist.appendChild(tab));
    this.prepend(tablist);
    this.tabs.forEach((tab, i) => {
      const key = tab.dataset.tab ?? String(i);
      tab.id = `${this.id}-tab-${key}`;
      tab.setAttribute("role", "tab");
      tab.setAttribute("aria-controls", `${this.id}-panel-${key}`);
      tab.addEventListener("click", () => this.select(key));
      tab.addEventListener("keydown", this.handleKeydown);
      const panel = this.panels.find((p) => p.dataset.tab === key);
      if (panel) {
        panel.id = `${this.id}-panel-${key}`;
        panel.setAttribute("role", "tabpanel");
        panel.setAttribute("aria-labelledby", tab.id);
        panel.tabIndex = 0;
      }
    });
    if (this.syncHash) window.addEventListener("hashchange", this.handleHashChange);
    const hashKey = location.hash.slice(1);
    const initial = this.syncHash && this.tabs.some((t) => t.dataset.tab === hashKey) ? hashKey : this.tabs[0].dataset.tab ?? "";
    this.render(initial);
  }
  disconnectedCallback() {
    if (this.syncHash) window.removeEventListener("hashchange", this.handleHashChange);
  }
  get syncHash() {
    return this.hasAttribute("sync-hash");
  }
  select(key) {
    if (this.syncHash) {
      location.hash = key;
    } else {
      this.render(key);
    }
  }
  render(key) {
    this.tabs.forEach((tab) => {
      const selected = tab.dataset.tab === key;
      tab.setAttribute("aria-selected", String(selected));
      tab.tabIndex = selected ? 0 : -1;
    });
    this.panels.forEach((panel) => {
      panel.hidden = panel.dataset.tab !== key;
    });
  }
  handleHashChange = () => {
    const key = location.hash.slice(1);
    if (this.tabs.some((t) => t.dataset.tab === key)) this.render(key);
  };
  handleKeydown = (event) => {
    const currentIndex = this.tabs.findIndex((t) => t.id === event.target.id);
    if (currentIndex === -1) return;
    let nextIndex;
    switch (event.key) {
      case "ArrowRight":
        nextIndex = (currentIndex + 1) % this.tabs.length;
        break;
      case "ArrowLeft":
        nextIndex = (currentIndex - 1 + this.tabs.length) % this.tabs.length;
        break;
      case "Home":
        nextIndex = 0;
        break;
      case "End":
        nextIndex = this.tabs.length - 1;
        break;
      default:
        return;
    }
    event.preventDefault();
    const next = this.tabs[nextIndex];
    this.select(next.dataset.tab ?? "");
    next.focus();
  };
};
customElements.define("yr-tabs", YrTabs);

// pkg/author_link.ts
var YrAuthorLink = class extends HTMLElement {
};
customElements.define("yr-author-link", YrAuthorLink);

// pkg/nav_shell.ts
var YrNavShell = class extends HTMLElement {
};
var YrNavShellLeft = class extends HTMLElement {
};
var YrNavShellCenter = class extends HTMLElement {
};
var YrNavShellRight = class extends HTMLElement {
};
customElements.define("yr-nav-shell", YrNavShell);
customElements.define("yr-nav-shell-left", YrNavShellLeft);
customElements.define("yr-nav-shell-center", YrNavShellCenter);
customElements.define("yr-nav-shell-right", YrNavShellRight);

// pkg/navbar.ts
var YrNavbar = class extends HTMLElement {
};
var YrNavbarStd = class extends HTMLElement {
};
customElements.define("yr-navbar", YrNavbar);
customElements.define("yr-navbar-std", YrNavbarStd);

// pkg/sidebar.ts
var YrSidebar = class extends HTMLElement {
  static observedAttributes = ["collapsed"];
  attributeChangedCallback() {
  }
};
var YrSidebarHeader = class extends HTMLElement {
  connectedCallback() {
    const toggle = this.querySelector(".yr-sidebar-toggle");
    if (toggle) {
      toggle.addEventListener("click", () => {
        const sidebar = this.closest("yr-sidebar");
        if (sidebar) sidebar.toggleAttribute("collapsed");
      });
    }
  }
};
var YrSidebarNav = class extends HTMLElement {
};
var YrSidebarFooter = class extends HTMLElement {
};
var YrSidebarLink = class extends HTMLElement {
  static observedAttributes = ["href", "active"];
  onClick = (e) => {
    const href = this.getAttribute("href");
    if (!href) return;
    e.preventDefault();
    window.location.href = href;
  };
  connectedCallback() {
    this.setAttribute("role", "link");
    this.setAttribute("tabindex", "0");
    this.addEventListener("click", this.onClick);
    this.addEventListener("keydown", this.onKeyDown);
  }
  disconnectedCallback() {
    this.removeEventListener("click", this.onClick);
    this.removeEventListener("keydown", this.onKeyDown);
  }
  onKeyDown = (e) => {
    if (e.key === "Enter") this.click();
  };
  attributeChangedCallback() {
  }
};
var YrSidebarOrg = class extends HTMLElement {
};
customElements.define("yr-sidebar", YrSidebar);
customElements.define("yr-sidebar-header", YrSidebarHeader);
customElements.define("yr-sidebar-nav", YrSidebarNav);
customElements.define("yr-sidebar-footer", YrSidebarFooter);
customElements.define("yr-sidebar-link", YrSidebarLink);
customElements.define("yr-sidebar-org", YrSidebarOrg);

// pkg/drawer.ts
var YrDrawer = class extends HTMLElement {
  static observedAttributes = ["open"];
  onKeyDown = (e) => {
    if (e.key === "Escape" && this.hasAttribute("open")) {
      this.removeAttribute("open");
    }
  };
  connectedCallback() {
    document.addEventListener("keydown", this.onKeyDown);
  }
  disconnectedCallback() {
    document.removeEventListener("keydown", this.onKeyDown);
  }
  attributeChangedCallback() {
  }
};
var YrDrawerHeader = class extends HTMLElement {
};
var YrDrawerBody = class extends HTMLElement {
};
var YrDrawerClose = class extends HTMLElement {
  onClick = () => {
    this.closest("yr-drawer")?.removeAttribute("open");
  };
  connectedCallback() {
    this.setAttribute("role", "button");
    this.setAttribute("tabindex", "0");
    this.setAttribute("aria-label", "Close drawer");
    this.addEventListener("click", this.onClick);
    this.addEventListener("keydown", this.onKeyDown);
  }
  disconnectedCallback() {
    this.removeEventListener("click", this.onClick);
    this.removeEventListener("keydown", this.onKeyDown);
  }
  onKeyDown = (e) => {
    if (e.key === "Enter" || e.key === " ") {
      e.preventDefault();
      this.onClick();
    }
  };
};
customElements.define("yr-drawer", YrDrawer);
customElements.define("yr-drawer-header", YrDrawerHeader);
customElements.define("yr-drawer-body", YrDrawerBody);
customElements.define("yr-drawer-close", YrDrawerClose);

// pkg/content.ts
var YrContent = class extends HTMLElement {
};
var YrContentMain = class extends HTMLElement {
};
customElements.define("yr-content", YrContent);
customElements.define("yr-content-main", YrContentMain);

// pkg/resize_handle.ts
var YrResizeHandle = class extends HTMLElement {
  static observedAttributes = ["target"];
  startX = 0;
  startWidth = 0;
  get prop() {
    return this.getAttribute("target") === "drawer" ? "--yr-drawer-width" : "--yr-sidebar-width";
  }
  get isSidebar() {
    return this.getAttribute("target") !== "drawer";
  }
  resizeTarget() {
    const tag = this.isSidebar ? "yr-sidebar" : "yr-drawer";
    const shell = this.closest("body") ?? this.parentElement;
    return shell?.querySelector(tag);
  }
  onMouseMove = (e) => {
    e.preventDefault();
    const delta = this.isSidebar ? e.clientX - this.startX : this.startX - e.clientX;
    document.documentElement.style.setProperty(
      this.prop,
      `${Math.max(100, this.startWidth + delta)}px`
    );
  };
  onMouseUp = () => {
    this.removeAttribute("active");
    document.documentElement.classList.remove("yr-resizing");
    document.body.style.cursor = "";
    document.body.style.userSelect = "";
    document.removeEventListener("mousemove", this.onMouseMove);
    document.removeEventListener("mouseup", this.onMouseUp);
  };
  onMouseDown = (e) => {
    e.preventDefault();
    const target = this.resizeTarget();
    if (!target) return;
    this.startX = e.clientX;
    this.startWidth = target.getBoundingClientRect().width;
    this.setAttribute("active", "");
    document.documentElement.classList.add("yr-resizing");
    document.body.style.cursor = "col-resize";
    document.body.style.userSelect = "none";
    document.addEventListener("mousemove", this.onMouseMove);
    document.addEventListener("mouseup", this.onMouseUp);
  };
  connectedCallback() {
    this.addEventListener("mousedown", this.onMouseDown);
  }
  disconnectedCallback() {
    this.removeEventListener("mousedown", this.onMouseDown);
    document.removeEventListener("mousemove", this.onMouseMove);
    document.removeEventListener("mouseup", this.onMouseUp);
  }
  attributeChangedCallback() {
  }
};
customElements.define("yr-resize-handle", YrResizeHandle);

// pkg/page_workspace.ts
var YrPageWorkspace = class extends HTMLElement {
  connectedCallback() {
    this.setAttribute("role", "main");
  }
};
customElements.define("yr-page-workspace", YrPageWorkspace);

// pkg/page_auth.ts
var YrPageAuth = class extends HTMLElement {
  connectedCallback() {
    this.setAttribute("role", "main");
  }
};
customElements.define("yr-page-auth", YrPageAuth);

// pkg/page_setup.ts
var YrPageSetup = class extends HTMLElement {
  connectedCallback() {
    this.setAttribute("role", "main");
  }
};
var YrPageSetupContent = class extends HTMLElement {
};
customElements.define("yr-page-setup", YrPageSetup);
customElements.define("yr-page-setup-content", YrPageSetupContent);

// pkg/setup_card.ts
var YrSetupCard = class extends HTMLElement {
};
var YrSetupCardProgress = class extends HTMLElement {
};
var YrSetupCardHeader = class extends HTMLElement {
};
var YrSetupCardBody = class extends HTMLElement {
};
var YrSetupCardFooter = class extends HTMLElement {
};
customElements.define("yr-setup-card", YrSetupCard);
customElements.define("yr-setup-card-progress", YrSetupCardProgress);
customElements.define("yr-setup-card-header", YrSetupCardHeader);
customElements.define("yr-setup-card-body", YrSetupCardBody);
customElements.define("yr-setup-card-footer", YrSetupCardFooter);

// pkg/section_header.ts
var YrSectionHeader = class extends HTMLElement {
};
customElements.define("yr-section-header", YrSectionHeader);

// pkg/timeline_group.ts
var YrTimelineGroup = class extends HTMLElement {
};
var YrTimelineEntries = class extends HTMLElement {
};
customElements.define("yr-timeline-group", YrTimelineGroup);
customElements.define("yr-timeline-entries", YrTimelineEntries);

// pkg/hitl_prompt.ts
var YrHitlPrompt = class extends HTMLElement {
  static observedAttributes = ["tone"];
  connectedCallback() {
    if (!this.getAttribute("role")) this.setAttribute("role", "alert");
  }
};
customElements.define("yr-hitl-prompt", YrHitlPrompt);

// pkg/journal_log.ts
var YrJournalLog = class extends HTMLElement {
  connectedCallback() {
    if (!this.getAttribute("role")) this.setAttribute("role", "log");
    if (!this.getAttribute("aria-label")) this.setAttribute("aria-label", "Run journal");
  }
};
customElements.define("yr-journal-log", YrJournalLog);

// pkg/workflow_card.ts
var YrWorkflowCard = class extends HTMLElement {
};
customElements.define("yr-workflow-card", YrWorkflowCard);

// pkg/status_banner.ts
var YrStatusBanner = class extends HTMLElement {
  static observedAttributes = ["tone"];
};
customElements.define("yr-status-banner", YrStatusBanner);

// pkg/download_list.ts
var YrDownloadList = class extends HTMLElement {
};
customElements.define("yr-download-list", YrDownloadList);
//# sourceMappingURL=ui.js.map
