export type Player = {
	id: number;
	name: string;
	isUser?: boolean;
	isHost?: boolean;
	isComputer?: boolean;
	colour: Colour;
};

export const colours = [
	'red',
	'orange',
	'yellow',
	'lime',
	'emerald',
	'cyan',
	'blue',
	'indigo',
	'purple',
	'pink'
] as const;

export type Colour = (typeof colours)[number];
