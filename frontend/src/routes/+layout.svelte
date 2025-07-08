<script lang="ts">
	import '../app.css';
	import * as Sidebar from '$lib/components/ui/sidebar/index.js';
	import AppSidebar from '$lib/components/app-sidebar.svelte';
	import SiteHeader from '$lib/components/site-header.svelte';
	import SectionCards from '$lib/components/section-cards.svelte';
	import { ModeWatcher } from 'mode-watcher';

	let { children } = $props();
</script>

<ModeWatcher />
<Sidebar.Provider>
	<AppSidebar variant="inset" />
	<Sidebar.Inset>
		<SiteHeader />
		<div class="flex flex-1 flex-col">
			<div class="@container/main flex flex-col gap-4 py-4 md:gap-6 md:py-6">
				{@render children()}
			</div>
		</div>
	</Sidebar.Inset>
</Sidebar.Provider>
