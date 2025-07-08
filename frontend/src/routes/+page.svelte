<script lang="ts">
	import SelbstkostenChart from '$lib/components/selbstkosten-chart.svelte';
	import * as Card from '$lib/components/ui/card/index.js';
	import SectionCards from '$lib/components/section-cards.svelte';
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';

	let startingBin: bin = {
		NumberOfPointsInBin: 0,
		X1: 0,
		X0: 0
	};

	let bins: bin[] = $state([startingBin]);

	let defaultSimulationsanzahl = 100_000;
	let simulationsanzahl = $state(defaultSimulationsanzahl);
	let currentSimulationsanzahl = $state(defaultSimulationsanzahl);
	function handleSubmit(event: Event) {
		event.preventDefault();
		console.log(simulationsanzahl);

		fetch('http://127.0.0.1:4000/restapi/sim/run', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(simulationsanzahl) // Send the data from the store
		})
			.then(async (response) => {
				if (!response.ok) {
					throw new Error(`HTTP error! status: ${response.status}`);
				}
				response.json().then((jsonResponse) => {
					console.log(jsonResponse);
					bins = jsonResponse;
				});
			})
			.catch((err) => {
				console.error('Fetch error:', err);
				// TODO: add user feedback if error occurs
			});
		currentSimulationsanzahl = simulationsanzahl;
	}
</script>

<SectionCards />
<div class="px-4 lg:px-6">
	<Card.Root class="@container/card max-w-[100rem]">
		<Card.Header>
			<Card.Title>Stückkosteneinttrittswahrscheinlichkeit</Card.Title>
			<Card.Description>
				<span>Die Wahrscheinlichkeit dafür, dass die Stückkosten in diesem Bereich fallen</span>
			</Card.Description>
			<Card.Action></Card.Action>
		</Card.Header>
		<Card.Content>
			<SelbstkostenChart providedSimulationsanzahl={currentSimulationsanzahl} {bins} />
		</Card.Content>
	</Card.Root>
	<Card.Root class="@container/car my-10 ">
		<Card.Header>
			<Card.Title>Simulationseinstellungen</Card.Title>
			<Card.Description>
				<span>Einstellungen die die Simulations beeinflussen</span>
			</Card.Description>
			<Card.Action></Card.Action>
		</Card.Header>
		<Card.Content>
			<form onsubmit={handleSubmit}>
				<label class="">
					Simulationsanzahl
					<Input class="rounded-md " type="number" bind:value={simulationsanzahl} />
				</label>

				<Button class="mt-1 mb-2" type="submit">Starten</Button>
			</form>
		</Card.Content>
	</Card.Root>
</div>
