<script lang="ts">
	import PlayersList from '$lib/components/players/players-list.svelte';
	import ShareLink from '$lib/components/share-link.svelte';
	import { Separator } from '$lib/components/ui/separator';
	import {
		ArrowUpNarrowWide,
		Flame,
		Layers3,
		Link2,
		PlayCircle,
		RefreshCw,
		Scale,
		Sparkle
	} from 'lucide-svelte';
	import { page } from '$app/stores';
	import Rule from '$lib/components/rule.svelte';
	import { SpecialCardsToggleGroup, type PlayingCard } from '$lib/components/card';
	import { Button } from '$lib/components/ui/button';

	let selectedCards: PlayingCard[] = [];
</script>

<div class="flex flex-1 flex-grow flex-row gap-8 py-5">
	<PlayersList class="w-1/5" />
	<Separator orientation="vertical" class="h-auto" />
	<div class="flex-1 space-y-5">
		<div class="space-y-3">
			<h3 class="flex items-center gap-2 text-lg font-medium text-primary">
				<span>Invite Your Friends</span><Link2 />
			</h3>
			<div class="flex gap-2">
				<ShareLink link={'https://president.basokant.com' + $page.url.pathname} />
				<Button>
					<PlayCircle />
					<span>Start Game</span>
				</Button>
			</div>
		</div>
		<div class="space-y-2">
			<h3 class="flex items-center gap-2 text-lg font-medium text-primary">
				<span>Custom Rules</span><Scale />
			</h3>
			<div class="grid grid-cols-2 gap-x-5">
				<Rule
					id="communism"
					name="Communism Trades"
					description="Reverse the trades, so that the President trades their two highest rank cards with the Bum’s two lowest rank cards."
				>
					<RefreshCw />
				</Rule>
				<Rule
					id="duplicate"
					name="Duplicate Plays Burn"
					description="When a duplicate play is made, the pile is burned, and the last player can start the pile again."
				>
					<Flame />
				</Rule>
				<Rule
					id="bum-start"
					name="Bum Starts the Game"
					description="The Bum player starts first, instead of always starting with whoever has the 3 of ♣’s"
				>
					<ArrowUpNarrowWide />
				</Rule>
				<Rule
					id="runs"
					name="Card Runs"
					description="Players are able to start a pile with a play of consecutive runs of cards (up to 4) in ascending rank."
				>
					<Layers3 />
				</Rule>
			</div>
		</div>
		<div class="space-y-3">
			<h3 class="flex items-center gap-2 text-lg font-medium text-primary">
				<span>Special Cards</span><Sparkle />
			</h3>
			<SpecialCardsToggleGroup bind:selectedCards />
		</div>
	</div>
</div>
