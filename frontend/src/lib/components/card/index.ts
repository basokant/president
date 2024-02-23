import Root from './card.svelte';
import SpecialCardsToggleGroup from './special-cards-toggle-group.svelte';

export const suits = ['hearts', 'clubs', 'diamonds', 'spades'] as const;
export type Suit = (typeof suits)[number];

export const rank = [
	'3',
	'4',
	'5',
	'6',
	'7',
	'8',
	'9',
	'10',
	'jack',
	'queen',
	'king',
	'ace',
	'2'
] as const;
export type Rank = (typeof rank)[number];

export type PlayingCard =
	| {
			rank: Rank;
			suit: Suit;
			name?: string;
	  }
	| 'joker';
export type CardId = `${Rank}-of-${Suit}` | 'joker-1' | 'joker-2';

export function getCardId(card: PlayingCard): CardId {
	const isJoker = card == 'joker';
	if (isJoker) {
		return 'joker-1';
	}

	return `${card.rank}-of-${card.suit}`;
}

export const sizes = {
	sm: '100px',
	md: '165px',
	lg: '215px',
	xl: '300px'
} as const;

export { Root as Card, SpecialCardsToggleGroup };
