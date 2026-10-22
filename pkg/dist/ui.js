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
    const shell = this.closest("body") ?? this.parentElement;
    if (this.isSidebar) return shell?.querySelector("yr-sidebar");
    return shell?.querySelector("yr-drawer[open], yr-cappy-drawer[open]") ?? shell?.querySelector("yr-drawer, yr-cappy-drawer");
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

// pkg/chat_fab.ts
var MARGIN = 24;
var SIZE = 56;
var YrChatFab = class extends HTMLElement {
  startX = 0;
  startY = 0;
  prevX = 0;
  prevY = 0;
  prevTime = 0;
  velX = 0;
  velY = 0;
  isDragging = false;
  hasMoved = false;
  connectedCallback() {
    this.trigger?.addEventListener("mousedown", this.onPointerDown);
    this.trigger?.addEventListener("click", this.onClick);
    this.querySelector(".yr-chat-fab__dismiss")?.addEventListener("click", this.onDismiss);
  }
  disconnectedCallback() {
    this.trigger?.removeEventListener("mousedown", this.onPointerDown);
    this.trigger?.removeEventListener("click", this.onClick);
    this.querySelector(".yr-chat-fab__dismiss")?.removeEventListener("click", this.onDismiss);
    document.removeEventListener("mousemove", this.onPointerMove);
    document.removeEventListener("mouseup", this.onPointerUp);
  }
  get trigger() {
    return this.querySelector(".yr-chat-fab__trigger");
  }
  onClick = (e) => {
    if (this.hasMoved) {
      e.preventDefault();
      e.stopPropagation();
      return;
    }
    this.dispatchEvent(new CustomEvent("yr-chat-toggle", { bubbles: true }));
  };
  onDismiss = (e) => {
    e.stopPropagation();
    this.hidden = true;
    this.dispatchEvent(new CustomEvent("yr-chat-dismiss", { bubbles: true }));
  };
  onPointerDown = (e) => {
    this.startX = e.clientX;
    this.startY = e.clientY;
    this.prevX = e.clientX;
    this.prevY = e.clientY;
    this.prevTime = performance.now();
    this.velX = 0;
    this.velY = 0;
    this.isDragging = true;
    this.hasMoved = false;
    document.addEventListener("mousemove", this.onPointerMove);
    document.addEventListener("mouseup", this.onPointerUp);
  };
  onPointerMove = (e) => {
    if (!this.isDragging) return;
    const dx = e.clientX - this.startX;
    const dy = e.clientY - this.startY;
    if (!this.hasMoved && Math.abs(dx) < 5 && Math.abs(dy) < 5) return;
    if (!this.hasMoved) {
      this.hasMoved = true;
      this.setAttribute("dragging", "");
      this.removeAttribute("corner");
      this.clearPositionProps();
    }
    const now = performance.now();
    const dt = now - this.prevTime;
    if (dt > 0) {
      this.velX = (e.clientX - this.prevX) / dt;
      this.velY = (e.clientY - this.prevY) / dt;
    }
    this.prevX = e.clientX;
    this.prevY = e.clientY;
    this.prevTime = now;
    this.style.left = `${e.clientX - SIZE / 2}px`;
    this.style.top = `${e.clientY - SIZE / 2}px`;
  };
  onPointerUp = () => {
    document.removeEventListener("mousemove", this.onPointerMove);
    document.removeEventListener("mouseup", this.onPointerUp);
    this.isDragging = false;
    if (!this.hasMoved) return;
    const rect = this.getBoundingClientRect();
    const MOMENTUM = 300;
    const cx = rect.left + rect.width / 2 + this.velX * MOMENTUM;
    const cy = rect.top + rect.height / 2 + this.velY * MOMENTUM;
    const midX = window.innerWidth / 2;
    const midY = window.innerHeight / 2;
    const corner = cy < midY ? cx < midX ? "top-left" : "top-right" : cx < midX ? "bottom-left" : "bottom-right";
    const target = this.cornerPosition(corner);
    this.removeAttribute("dragging");
    this.style.top = `${target.top}px`;
    this.style.left = `${target.left}px`;
    const onDone = () => {
      this.removeEventListener("transitionend", onDone);
      this.clearPositionProps();
      this.setAttribute("corner", corner);
    };
    this.addEventListener("transitionend", onDone);
    setTimeout(onDone, 450);
  };
  cornerPosition(corner) {
    const w = window.innerWidth;
    const h = window.innerHeight;
    switch (corner) {
      case "top-left":
        return { top: MARGIN, left: MARGIN };
      case "top-right":
        return { top: MARGIN, left: w - SIZE - MARGIN };
      case "bottom-left":
        return { top: h - SIZE - MARGIN, left: MARGIN };
      case "bottom-right":
        return { top: h - SIZE - MARGIN, left: w - SIZE - MARGIN };
    }
  }
  clearPositionProps() {
    this.style.top = "";
    this.style.bottom = "";
    this.style.left = "";
    this.style.right = "";
  }
};
customElements.define("yr-chat-fab", YrChatFab);

// pkg/cappy_drawer.ts
var YrCappyDrawer = class extends HTMLElement {
  static observedAttributes = ["open"];
  onKeyDown = (e) => {
    if (e.key === "Escape" && this.hasAttribute("open")) {
      this.removeAttribute("open");
    }
  };
  connectedCallback() {
    document.addEventListener("keydown", this.onKeyDown);
    this.querySelector(".yr-cappy-drawer__close")?.addEventListener("click", this.close);
    this.querySelector(".yr-cappy-drawer__collapse")?.addEventListener("click", this.toggleCollapse);
    this.querySelector(".yr-cappy-drawer__escalate")?.addEventListener("click", this.escalate);
    this.querySelector(".yr-cappy-drawer__log")?.addEventListener("click", this.onLogClick);
  }
  disconnectedCallback() {
    document.removeEventListener("keydown", this.onKeyDown);
    this.querySelector(".yr-cappy-drawer__close")?.removeEventListener("click", this.close);
    this.querySelector(".yr-cappy-drawer__collapse")?.removeEventListener("click", this.toggleCollapse);
    this.querySelector(".yr-cappy-drawer__escalate")?.removeEventListener("click", this.escalate);
    this.querySelector(".yr-cappy-drawer__log")?.removeEventListener("click", this.onLogClick);
  }
  attributeChangedCallback() {
  }
  close = () => {
    this.removeAttribute("open");
  };
  toggleCollapse = () => {
    this.toggleAttribute("collapsed");
  };
  escalate = () => {
    this.dispatchEvent(new CustomEvent("yr-chat-escalate", { bubbles: true }));
  };
  onLogClick = (e) => {
    const msg = e.target.closest(".yr-chat-msg");
    if (!msg) return;
    msg.toggleAttribute("data-show-time");
  };
};
customElements.define("yr-cappy-drawer", YrCappyDrawer);

// pkg/chat_input.ts
var YrChatInput = class extends HTMLElement {
  staged = [];
  connectedCallback() {
    this.sendBtn?.addEventListener("click", this.handleSend);
    this.field?.addEventListener("keydown", this.handleKeydown);
    this.field?.addEventListener("input", this.handleAutoGrow);
    this.querySelector(".yr-chat-input__attach")?.addEventListener("click", this.handleAttach);
    this.fileInput?.addEventListener("change", this.handleFiles);
    this.stagingEl?.addEventListener("click", this.handleStagingClick);
  }
  disconnectedCallback() {
    this.sendBtn?.removeEventListener("click", this.handleSend);
    this.field?.removeEventListener("keydown", this.handleKeydown);
    this.field?.removeEventListener("input", this.handleAutoGrow);
    this.querySelector(".yr-chat-input__attach")?.removeEventListener("click", this.handleAttach);
    this.fileInput?.removeEventListener("change", this.handleFiles);
    this.stagingEl?.removeEventListener("click", this.handleStagingClick);
  }
  get field() {
    return this.querySelector(".yr-chat-input__field");
  }
  get sendBtn() {
    return this.querySelector(".yr-chat-input__send");
  }
  get fileInput() {
    return this.querySelector(".yr-chat-input__file");
  }
  get stagingEl() {
    return this.querySelector(".yr-chat-input__staging");
  }
  handleAutoGrow = () => {
    const el = this.field;
    if (!el) return;
    el.style.height = "auto";
    el.style.height = `${el.scrollHeight}px`;
  };
  handleSend = () => {
    const value = this.field?.value.trim();
    if (!value && this.staged.length === 0) return;
    this.dispatchEvent(new CustomEvent("yr-chat-send", {
      bubbles: true,
      detail: { text: value || "", files: [...this.staged] }
    }));
    if (this.field) {
      this.field.value = "";
      this.field.style.height = "auto";
    }
    this.staged = [];
    this.renderStaging();
  };
  handleKeydown = (e) => {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault();
      this.handleSend();
    }
  };
  handleAttach = () => {
    this.fileInput?.click();
  };
  handleFiles = () => {
    const files = this.fileInput?.files;
    if (!files) return;
    for (let i = 0; i < files.length; i++) {
      this.staged.push(files[i]);
    }
    if (this.fileInput) this.fileInput.value = "";
    this.renderStaging();
  };
  handleStagingClick = (e) => {
    const btn = e.target.closest(".yr-chat-input__chip-remove");
    if (!btn) return;
    const idx = Number(btn.getAttribute("data-index"));
    if (!Number.isNaN(idx)) {
      this.staged.splice(idx, 1);
      this.renderStaging();
    }
  };
  renderStaging() {
    const el = this.stagingEl;
    if (!el) return;
    if (this.staged.length === 0) {
      el.hidden = true;
      el.innerHTML = "";
      return;
    }
    el.hidden = false;
    el.innerHTML = this.staged.map((f, i) => this.chipHTML(f, i)).join("");
  }
  chipHTML(file, index) {
    const name = file.name.length > 20 ? file.name.slice(0, 17) + "..." : file.name;
    const isImage = file.type.startsWith("image/");
    const isAudio = file.type.startsWith("audio/");
    const icon = isImage ? `<svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><circle cx="9" cy="9" r="2"/><path d="m21 15-3.086-3.086a2 2 0 0 0-2.828 0L6 21"/></svg>` : isAudio ? `<svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M17.5 22h.5a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v3"/><path d="M14 2v5h5"/><path d="M2 19a2 2 0 1 1 4 0v1a2 2 0 1 1-4 0v-4a6 6 0 0 1 12 0v4a2 2 0 1 1-4 0v-1a2 2 0 1 1 4 0"/></svg>` : `<svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"/><path d="M14 2v4a2 2 0 0 0 2 2h4"/></svg>`;
    const removeIcon = `<svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>`;
    return `<div class="yr-chat-input__chip"><span class="yr-chat-input__chip-icon">${icon}</span><span class="yr-chat-input__chip-name">${this.escapeHTML(name)}</span><button type="button" class="yr-chat-input__chip-remove" data-index="${index}" aria-label="Remove ${this.escapeHTML(name)}">${removeIcon}</button></div>`;
  }
  escapeHTML(s) {
    const el = document.createElement("span");
    el.textContent = s;
    return el.innerHTML;
  }
};
customElements.define("yr-chat-input", YrChatInput);
//# sourceMappingURL=ui.js.map
