import Root from './card.svelte';

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
	  }
	| 'joker';
export type CardName = `${Rank}-of-${Suit}` | 'joker-1' | 'joker-2';

export function getCardName(card: PlayingCard): CardName {
	const isJoker = card == 'joker';
	if (isJoker) {
		return 'joker-1';
	}

	return `${card.rank}-of-${card.suit}`;
}

export { Root as Card };
