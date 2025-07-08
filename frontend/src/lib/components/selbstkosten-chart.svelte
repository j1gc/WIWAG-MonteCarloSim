<script lang="ts">
	import { scaleLinear, max, select, axisLeft, axisBottom, formatLocale } from 'd3';

	let { providedSimulationsanzahl, bins }: { providedSimulationsanzahl: number; bins: bin[] } =
		$props();

	let width = $state(0);

	const height = 400;
	const margin = { top: 30, right: 30, bottom: 30, left: 50 };

	let axisXEl: SVGGElement | undefined = $state();
	let axisYEl: SVGGElement | undefined = $state();

	let locale = formatLocale({
		decimal: ',',
		thousands: '.',
		grouping: [3],
		currency: ['', '\u00a0€']
	});
	locale.format(',.0f');

	let x = $derived(
		scaleLinear()
			.domain([bins[0].X0 ?? 0, bins[bins.length - 1].X1 ?? 0])
			.range([margin.left, width - margin.right])
	);

	let y = $derived(
		scaleLinear()
			.domain([0, max(bins, (d) => d.NumberOfPointsInBin / providedSimulationsanzahl) ?? 1])
			.range([height - margin.bottom, margin.top])
	);

	$effect(() => {
		// Ensure the element is mounted and the scale is calculated before drawing
		if (axisXEl && x && width > 0) {
			console.log('Updating X Axis, width:', width, 'domain:', x.domain());
			const xAxisGenerator = axisBottom(x)
				.tickFormat((domainValue) => {
					return locale.format(',.2f')(domainValue);
				})
				.ticks(width / 50) // Calculate ticks dynamically
				.tickSizeOuter(0);
			select(axisXEl).transition().duration(100).call(xAxisGenerator); // Add transition
		} else {
			// Optionally clear the axis if conditions aren't met
			if (axisXEl) select(axisXEl).selectAll('*').remove();
		}
		if (axisYEl && y) {
			console.log('Updating Y Axis, domain:', y.domain());
			const yAxisGenerator = axisLeft(y)
				.tickFormat((domainValue) => {
					return locale.format(',.0%')(domainValue);
				})
				.ticks(Math.max(3, height / 40)); // Ensure at least a few ticks
			select(axisYEl).transition().duration(200).call(yAxisGenerator); // Add transition
		} else {
			// Optionally clear the axis if conditions aren't met
			if (axisYEl) select(axisYEl).selectAll('*').remove();
		}
	});

	let showTooltip = $state(false);
	interface TooltipData {
		BarFrequency: number;
		BarX0: number;
		BarX1: number;
	}

	let tooltipData: TooltipData | undefined = $state();
	let tooltipEvent: MouseEvent | undefined = $state();
</script>

{#if tooltipData && tooltipEvent}
	<div
		class="absolute z-30 grid grid-cols-2 gap-x-5 rounded-sm bg-white p-1 text-xs text-gray-700"
		style="top: {tooltipEvent.clientY}px; left: {tooltipEvent.clientX}px; visibility: {showTooltip
			? 'visible'
			: 'hidden'};"
		id="tooltip"
	>
		<p>Simulationen</p>
		<p>
			{locale.format(',.1%')(tooltipData.BarFrequency / providedSimulationsanzahl || 0)}
		</p>
		<p>Min</p>
		<p>
			{tooltipData.BarX0.toLocaleString('DE')}
		</p>
		<p>Max</p>
		<p>
			{tooltipData.BarX1.toLocaleString('DE')}
		</p>
		<p>Durchschnitt</p>
		<p>{((tooltipData.BarX0 + tooltipData.BarX1) / 2).toLocaleString('DE')}</p>
	</div>
{/if}
<div bind:clientWidth={width}>
	<svg viewBox={`${0} ${0} ${width} ${height}`}>
		<g>
			{#if bins}
				{#each bins as bin}
					{#if bin.X0 && bin.X1}
						<rect
							onmousemove={(event: MouseEvent) => {
								showTooltip = true;
								tooltipData = {
									BarFrequency: bin.NumberOfPointsInBin,
									BarX0: bin.X0 ?? 0,
									BarX1: bin.X1 ?? 0
								};
								tooltipEvent = event;
							}}
							onmouseout={() => {
								showTooltip = false;
							}}
							onfocus={() => {}}
							onblur={() => {}}
							role="contentinfo"
							data-number={bin.NumberOfPointsInBin}
							data-x0={bin.X0}
							data-x1={bin.X1}
							x={x(bin.X0) + 1}
							y={y(bin.NumberOfPointsInBin / providedSimulationsanzahl)}
							width={x(bin.X1) - x(bin.X0) - 1}
							height={y(0) - y(bin.NumberOfPointsInBin / providedSimulationsanzahl)}
							fill="#0b3b87"
						></rect>
						<!--stroke="#0b3b87"-->
					{/if}
				{/each}
			{/if}
		</g>
		<g transform={`translate(0,${height - margin.bottom})`} bind:this={axisXEl}> </g>
		<g transform={`translate(${margin.left},0)`} bind:this={axisYEl}> </g>
	</svg>
</div>

<style>
	#tooltip {
		pointer-events: none;
		transition: 75ms ease;
	}

	rect {
		transition: all 550ms ease;
	}
</style>
