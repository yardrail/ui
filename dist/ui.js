// src/button/yr-button.ts
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

// src/pill/yr-pill.ts
var YrPill = class extends HTMLElement {
  static observedAttributes = ["tone", "shape"];
};
customElements.define("yr-pill", YrPill);

// src/avatar/yr-avatar.ts
var YrAvatar = class extends HTMLElement {
  static observedAttributes = ["size"];
};
customElements.define("yr-avatar", YrAvatar);

// src/alert/yr-alert.ts
var YrAlert = class extends HTMLElement {
  static observedAttributes = ["tone"];
};
customElements.define("yr-alert", YrAlert);

// src/stat/yr-stat.ts
var YrStat = class extends HTMLElement {
  static observedAttributes = ["number", "label"];
  connectedCallback() {
    this.render();
  }
  attributeChangedCallback() {
    this.render();
  }
  render() {
    this.replaceChildren();
    const numberEl = document.createElement("span");
    numberEl.className = "yr-stat-number";
    numberEl.textContent = this.getAttribute("number") ?? "";
    const labelEl = document.createElement("span");
    labelEl.className = "yr-stat-label";
    labelEl.textContent = this.getAttribute("label") ?? "";
    this.append(numberEl, labelEl);
  }
};
customElements.define("yr-stat", YrStat);

// src/rating/yr-rating.ts
var STAR_PATH = "M11.525 2.295a.53.53 0 0 1 .95 0l2.31 4.679a2.123 2.123 0 0 0 1.595 1.16l5.166.756a.53.53 0 0 1 .294.904l-3.736 3.638a2.123 2.123 0 0 0-.611 1.878l.882 5.14a.53.53 0 0 1-.771.56l-4.618-2.428a2.122 2.122 0 0 0-1.973 0L6.396 21.01a.53.53 0 0 1-.77-.56l.881-5.139a2.122 2.122 0 0 0-.611-1.879L2.16 9.795a.53.53 0 0 1 .294-.906l5.165-.755a2.122 2.122 0 0 0 1.597-1.16z";
function starSVG(filled) {
  const ns = "http://www.w3.org/2000/svg";
  const svg = document.createElementNS(ns, "svg");
  svg.setAttribute("width", "16");
  svg.setAttribute("height", "16");
  svg.setAttribute("viewBox", "0 0 24 24");
  svg.setAttribute("fill", filled ? "currentColor" : "none");
  svg.setAttribute("stroke", "currentColor");
  svg.setAttribute("stroke-width", "2");
  svg.setAttribute("stroke-linecap", "round");
  svg.setAttribute("stroke-linejoin", "round");
  const path = document.createElementNS(ns, "path");
  path.setAttribute("d", STAR_PATH);
  svg.appendChild(path);
  return svg;
}
var YrRating = class extends HTMLElement {
  static observedAttributes = ["value", "reviews"];
  connectedCallback() {
    this.render();
  }
  attributeChangedCallback() {
    this.render();
  }
  render() {
    const raw = parseFloat(this.getAttribute("value") ?? "0");
    const filled = Math.max(0, Math.min(5, Math.round(raw)));
    const reviews = this.getAttribute("reviews");
    this.replaceChildren();
    const stars = document.createElement("span");
    stars.className = "yr-rating-stars";
    for (let i = 0; i < 5; i++) {
      stars.appendChild(starSVG(i < filled));
    }
    this.append(stars);
    if (reviews) {
      const count = document.createElement("span");
      count.className = "yr-rating-count";
      count.textContent = `(${reviews} reviews)`;
      this.append(count);
    }
  }
};
customElements.define("yr-rating", YrRating);

// src/empty-state/yr-empty-state.ts
var YrEmptyState = class extends HTMLElement {
};
customElements.define("yr-empty-state", YrEmptyState);

// src/copy-field/yr-copy-field.ts
var YrCopyField = class extends HTMLElement {
  static observedAttributes = ["value"];
  resetTimer;
  connectedCallback() {
    this.render();
  }
  attributeChangedCallback() {
    this.render();
  }
  disconnectedCallback() {
    if (this.resetTimer !== void 0) window.clearTimeout(this.resetTimer);
  }
  render() {
    const value = this.getAttribute("value") ?? "";
    this.replaceChildren();
    const code = document.createElement("code");
    code.className = "yr-copy-field-code";
    code.textContent = value;
    const button = document.createElement("button");
    button.type = "button";
    button.className = "yr-copy-field-button";
    button.textContent = "Copy";
    button.addEventListener("click", () => this.handleCopy(value, button));
    this.append(code, button);
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

// src/dropdown-menu/yr-dropdown-menu.ts
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

// src/account-menu/yr-account-menu.ts
var YrAccountMenu = class extends YrDropdownMenu {
  connectedCallback() {
    this.restructure();
    super.connectedCallback();
  }
  restructure() {
    if (this.querySelector(':scope > [slot="trigger"]')) return;
    const name = this.getAttribute("name") ?? "";
    const email = this.getAttribute("email") ?? "";
    const avatarSlot = this.querySelector('[slot="avatar"]');
    const triggerExtra = this.querySelector('[slot="trigger-extra"]');
    const items = Array.from(this.childNodes).filter((node) => node !== avatarSlot && node !== triggerExtra);
    const trigger = document.createElement("button");
    trigger.setAttribute("slot", "trigger");
    trigger.className = "yr-account-menu-trigger";
    if (triggerExtra) trigger.append(triggerExtra.cloneNode(true));
    const triggerAvatar = document.createElement("yr-avatar");
    if (avatarSlot) triggerAvatar.append(avatarSlot.cloneNode(true));
    trigger.append(triggerAvatar, this.buildChevron());
    const header = document.createElement("div");
    header.className = "yr-dropdown-menu-header";
    const headerAvatar = document.createElement("yr-avatar");
    headerAvatar.setAttribute("size", "lg");
    if (avatarSlot) headerAvatar.append(avatarSlot.cloneNode(true));
    header.append(headerAvatar);
    if (name || email) {
      const headerText = document.createElement("div");
      if (name) {
        const nameEl = document.createElement("div");
        nameEl.className = "yr-dropdown-menu-header-name";
        nameEl.textContent = name;
        headerText.append(nameEl);
      }
      if (email) {
        const emailEl = document.createElement("div");
        emailEl.className = "yr-dropdown-menu-header-email";
        emailEl.textContent = email;
        headerText.append(emailEl);
      }
      header.append(headerText);
    }
    const panel = document.createElement("div");
    panel.setAttribute("slot", "panel");
    panel.append(header, ...items);
    this.replaceChildren(trigger, panel);
  }
  buildChevron() {
    const svg = document.createElementNS("http://www.w3.org/2000/svg", "svg");
    svg.setAttribute("class", "yr-dropdown-menu-chevron");
    svg.setAttribute("width", "12");
    svg.setAttribute("height", "12");
    svg.setAttribute("viewBox", "0 0 24 24");
    svg.setAttribute("fill", "none");
    svg.setAttribute("stroke", "currentColor");
    svg.setAttribute("stroke-width", "2.5");
    svg.setAttribute("stroke-linecap", "round");
    svg.setAttribute("stroke-linejoin", "round");
    const poly = document.createElementNS("http://www.w3.org/2000/svg", "polyline");
    poly.setAttribute("points", "6 9 12 15 18 9");
    svg.append(poly);
    return svg;
  }
};
customElements.define("yr-account-menu", YrAccountMenu);

// src/live-duration/yr-live-duration.ts
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

// src/tabs/yr-tabs.ts
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
    let nextIndex = null;
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

// src/author-link/yr-author-link.ts
var VERIFIED_BADGE_PATH = "M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z";
var YrAuthorLink = class extends HTMLElement {
  static observedAttributes = ["href", "name", "handle", "verified", "avatar", "size", "heading"];
  connectedCallback() {
    this.render();
  }
  attributeChangedCallback() {
    this.render();
  }
  render() {
    this.replaceChildren();
    const name = this.getAttribute("name");
    const handle = this.getAttribute("handle");
    const verified = this.hasAttribute("verified");
    const avatarInitials = this.getAttribute("avatar");
    const href = this.getAttribute("href");
    const isLarge = this.getAttribute("size") === "lg";
    const inner = document.createElement(href ? "a" : "span");
    inner.className = "yr-author-link-inner";
    if (href) inner.href = href;
    if (avatarInitials) {
      const avatar = document.createElement("yr-avatar");
      if (isLarge) avatar.setAttribute("size", "lg");
      avatar.textContent = avatarInitials;
      inner.appendChild(avatar);
    }
    const badge = () => {
      const el = document.createElement("span");
      el.className = "yr-author-link-badge";
      el.title = "Verified author";
      const svgSize = isLarge ? 16 : 12;
      el.innerHTML = `<svg width="${svgSize}" height="${svgSize}" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><path d="${VERIFIED_BADGE_PATH}"></path></svg>`;
      return el;
    };
    const nameTag = this.hasAttribute("heading") ? "h1" : "strong";
    if (isLarge && name && handle) {
      const text = document.createElement("div");
      text.className = "yr-author-link-text";
      const nameRow = document.createElement("div");
      nameRow.className = "yr-author-link-name-row";
      const nameEl = document.createElement(nameTag);
      nameEl.className = "yr-author-link-name";
      nameEl.textContent = name;
      nameRow.appendChild(nameEl);
      if (verified) nameRow.appendChild(badge());
      text.appendChild(nameRow);
      const handleEl = document.createElement("p");
      handleEl.className = "yr-author-link-handle";
      handleEl.textContent = `@${handle}`;
      text.appendChild(handleEl);
      inner.appendChild(text);
    } else if (name) {
      const nameEl = document.createElement(nameTag);
      nameEl.className = "yr-author-link-name";
      nameEl.textContent = name;
      inner.appendChild(nameEl);
      if (verified) inner.appendChild(badge());
    } else if (handle) {
      inner.appendChild(document.createTextNode(`@${handle}`));
    }
    this.appendChild(inner);
  }
};
customElements.define("yr-author-link", YrAuthorLink);

// src/nav-shell/yr-nav-shell.ts
var YrNavShell = class extends HTMLElement {
  connectedCallback() {
    this.setAttribute("role", "navigation");
    this.restructure();
  }
  restructure() {
    if (this.querySelector("yr-nav-shell-right")) return;
    const left = this.querySelector('[slot="left"]');
    const center = this.querySelector('[slot="center"]');
    const leftWrapper = document.createElement("yr-nav-shell-left");
    if (left) leftWrapper.appendChild(left);
    const centerWrapper = document.createElement("yr-nav-shell-center");
    if (center) centerWrapper.appendChild(center);
    const rightWrapper = document.createElement("yr-nav-shell-right");
    const remaining = Array.from(this.childNodes).filter(
      (n) => n !== leftWrapper && n !== centerWrapper
    );
    for (const child of remaining) rightWrapper.appendChild(child);
    this.replaceChildren(leftWrapper, centerWrapper, rightWrapper);
  }
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

// src/sidebar/yr-sidebar.ts
var YrSidebar = class extends HTMLElement {
  static observedAttributes = ["collapsed"];
  attributeChangedCallback() {
  }
};
var YrSidebarHeader = class extends HTMLElement {
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
customElements.define("yr-sidebar", YrSidebar);
customElements.define("yr-sidebar-header", YrSidebarHeader);
customElements.define("yr-sidebar-nav", YrSidebarNav);
customElements.define("yr-sidebar-footer", YrSidebarFooter);
customElements.define("yr-sidebar-link", YrSidebarLink);

// src/drawer/yr-drawer.ts
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

// src/content/yr-content.ts
var YrContent = class extends HTMLElement {
};
var YrContentMain = class extends HTMLElement {
};
customElements.define("yr-content", YrContent);
customElements.define("yr-content-main", YrContentMain);

// src/resize-handle/yr-resize-handle.ts
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

// src/page-workspace/yr-page-workspace.ts
var YrPageWorkspace = class extends HTMLElement {
  connectedCallback() {
    this.setAttribute("role", "main");
  }
};
customElements.define("yr-page-workspace", YrPageWorkspace);
//# sourceMappingURL=ui.js.map
