<script lang="ts">
  import { Tabs } from "melt/builders";
  import BrowserShareTab from "$lib/components/filesend/BrowserShareTab.svelte";
  import FileSendTab from "$lib/components/filesend/FileSendTab.svelte";

  const tabNames = ["filesend", "browsershare"];
  const tabs = new Tabs({
    value: tabNames[0],
  });
</script>

<div class="w-full h-full grid place-items-center transition-all">
  <div
    class="card preset-filled-surface-100-900 border-[1px] border-surface-200-800 w-full max-w-lg"
  >
    <header
      {...tabs.triggerList}
      class="border-b-2 border-surface-200-800 flex flex-2 place-content-evenly"
    >
      <button
        class="p-2 pl-2 w-full"
        class:bg-primary-300-700={tabs.value == "filesend"}
        {...tabs.getTrigger(tabNames[0])}
      >
        <h4 class="h4">Send files</h4>
      </button>
      <button
        class="p-2 pl-2 w-full"
        class:bg-primary-300-700={tabs.value == "browsershare"}
        {...tabs.getTrigger(tabNames[1])}
      >
        <h4 class="h4">Browser share</h4>
      </button>
    </header>
    <div id="tabs-content" style="transition: all 0.4s;">
      <div {...tabs.getContent(tabNames[0])}>
        <FileSendTab />
      </div>
      <div {...tabs.getContent(tabNames[1])}>
        <BrowserShareTab />
      </div>
    </div>
  </div>
</div>
