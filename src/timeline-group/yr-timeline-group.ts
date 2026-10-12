// <yr-timeline-group aria-label="Today"><time datetime="2026-10-11">Today
// </time><yr-timeline-entries>...</yr-timeline-entries></yr-timeline-group>
// — a date-labeled container with a vertical time rail connecting child
// entries. Light DOM, styled by yr-timeline-group.css; no behavior.
export class YrTimelineGroup extends HTMLElement {}
export class YrTimelineEntries extends HTMLElement {}

customElements.define('yr-timeline-group', YrTimelineGroup);
customElements.define('yr-timeline-entries', YrTimelineEntries);
