<script lang="ts">
	import * as Sidebar from '$lib/components/ui/sidebar/index.js';
	import NavMain from '$lib/components/nav-main.svelte';
	import NavDocuments from './nav-documents.svelte';
	import type { ComponentProps } from 'svelte';
	import NavSecondary from './nav-secondary.svelte';
	const data = {
		user: {
			name: 'shadcn',
			email: 'm@example.com',
			avatar: '/avatars/shadcn.jpg'
		},
		navMain: [
			{
				title: 'Dashboard',
				url: '#'
			},
			{
				title: 'Lifecycle',
				url: '#'
			},
			{
				title: 'Analytics',
				url: '#'
			},
			{
				title: 'Projects',
				url: '#'
			},
			{
				title: 'Team',
				url: '#'
			}
		],
		navClouds: [
			{
				title: 'Capture',
				isActive: true,
				url: '#',
				items: [
					{
						title: 'Active Proposals',
						url: '#'
					},
					{
						title: 'Archived',
						url: '#'
					}
				]
			},
			{
				title: 'Proposal',
				url: '#',
				items: [
					{
						title: 'Active Proposals',
						url: '#'
					},
					{
						title: 'Archived',
						url: '#'
					}
				]
			},
			{
				title: 'Prompts',
				url: '#',
				items: [
					{
						title: 'Active Proposals',
						url: '#'
					},
					{
						title: 'Archived',
						url: '#'
					}
				]
			}
		],
		navSecondary: [
			{
				title: 'Settings',
				url: '/settings'
			},
			{
				title: 'Get Help',
				url: '#'
			},
			{
				title: 'Search',
				url: '#'
			}
		],
		documents: [
			{
				name: 'Data Library',
				url: '#'
			},
			{
				name: 'Reports',
				url: '#'
			},
			{
				name: 'Word Assistant',
				url: '#'
			}
		]
	};
	let { ...restProps }: ComponentProps<typeof Sidebar.Root> = $props();
</script>

<Sidebar.Root collapsible="offcanvas" {...restProps}>
	<Sidebar.Header>
		<Sidebar.Menu>
			<Sidebar.MenuItem>
				<Sidebar.MenuButton class="data-[slot=sidebar-menu-button]:!p-1.5">
					{#snippet child({ props })}
						<a href="##" {...props}>
							<div class="flex items-center gap-2">
								<img src="/AeroShoesLogo.png" alt="logo" class="!size-5" />
								<h1 class="text-base font-semibold">AeroShoes Foundry</h1>
							</div>
						</a>
					{/snippet}
				</Sidebar.MenuButton>
			</Sidebar.MenuItem>
		</Sidebar.Menu>
	</Sidebar.Header>
	<Sidebar.Content>
		<NavMain items={data.navMain} />
		<NavDocuments items={data.documents} />
		<NavSecondary items={data.navSecondary} />
	</Sidebar.Content>
	<Sidebar.Footer></Sidebar.Footer>
</Sidebar.Root>
