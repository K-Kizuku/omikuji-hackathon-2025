"use client";

import { useState, useEffect } from "react";

type Skill = {
	id: number;
	name: string;
	attack: number;
};

type Character = {
	name: string;
	hp: number;
	maxHp: number; // 最大HPを追加
	attack: number;
	defense: number;
	speed: number;
	skills: Skill[];
};

type Turn = "player" | "enemy";

export default function BattlePage() {
	const [player, setPlayer] = useState<Character>({
		name: "プレイヤー",
		hp: 150,
		maxHp: 150,
		attack: 50,
		defense: 30,
		speed: 25,
		skills: [
			{ id: 1, name: "火炎放射", attack: 40 },
			{ id: 2, name: "しっぽ攻撃", attack: 20 },
		],
	});

	const [enemy, setEnemy] = useState<Character>({
		name: "スライム",
		hp: 100,
		maxHp: 100,
		attack: 30,
		defense: 10,
		speed: 15,
		skills: [],
	});

	const [turn, setTurn] = useState<Turn>("player");
	const [battleLog, setBattleLog] = useState<string[]>([]);

	const calculateDamage = (
		attacker: Character,
		defender: Character,
	): number => {
		return Math.max(0, attacker.attack - defender.defense);
	};

	const handleAttack = () => {
		if (turn === "player") {
			const damage = calculateDamage(player, enemy);
			setEnemy((prev) => ({ ...prev, hp: Math.max(0, prev.hp - damage) }));
			setBattleLog((prev) => [
				...prev,
				`${player.name}の攻撃！ ${damage}のダメージ！`,
			]);
			setTurn("enemy");
		}
	};

	const enemyTurn = () => {
		const damage = calculateDamage(enemy, player);
		setPlayer((prev) => ({ ...prev, hp: Math.max(0, prev.hp - damage) }));
		setBattleLog((prev) => [
			...prev,
			`${enemy.name}の攻撃！ ${damage}のダメージ！`,
		]);
		setTurn("player");
	};

	useEffect(() => {
		if (turn === "enemy") {
			setTimeout(enemyTurn, 1000);
		}
	}, [turn, enemyTurn]);

	const handleSkill = (skill: Skill) => {
		if (turn === "player") {
			const damage = skill.attack;
			setEnemy((prev) => ({ ...prev, hp: Math.max(0, prev.hp - damage) }));
			setBattleLog((prev) => [
				...prev,
				`${player.name}は${skill.name}を使った！ ${damage}のダメージ！`,
			]);
			setTurn("enemy");
		}
	};

	return (
		<main className="flex flex-col items-center justify-between h-screen bg-gray-900 text-white">
			<h1 className="text-2xl font-bold mt-4">バトル画面</h1>

			<div className="flex justify-around w-full max-w-4xl mt-8">
				<div className="text-center">
					<p className="text-lg">{player.name}</p>
					<div className="relative w-32 h-4 bg-gray-700 rounded">
						<div
							style={{ width: `${(player.hp / player.maxHp) * 100}%` }}
							className="absolute h-full bg-green-500 rounded"
						/>
					</div>
					<p>
						{player.hp} / {player.maxHp}
					</p>
				</div>

				<div className="text-center">
					<p className="text-lg">{enemy.name}</p>
					<div className="relative w-32 h-4 bg-gray-700 rounded">
						<div
							style={{ width: `${(enemy.hp / enemy.maxHp) * 100}%` }}
							className="absolute h-full bg-red-500 rounded"
						/>
					</div>
					<p>
						{enemy.hp} / {enemy.maxHp}
					</p>
				</div>
			</div>

			{turn === "player" && (
				<div className="mt-8 w-full max-w-4xl">
					<table className="table-auto w-full bg-gray-800 rounded">
						<tbody>
							<tr>
								<td>攻撃</td>
								<td>
									<button
										type="button"
										className="bg-blue-500 px-4 py-2 rounded"
										onClick={handleAttack}
									>
										実行
									</button>
								</td>
							</tr>
							{player.skills.map((skill) => (
								<tr key={skill.id}>
									<td>{skill.name}</td>
									<td>
										<button
											type="button"
											className="bg-green-500 px-4 py-2 rounded"
											onClick={() => handleSkill(skill)}
										>
											実行
										</button>
									</td>
								</tr>
							))}
						</tbody>
					</table>
				</div>
			)}

			<div className="w-full max-w-4xl bg-gray-800 p-4 mt-4 rounded">
				<p className="font-bold">バトルログ</p>
				<div className="h-32 overflow-y-auto">
					{battleLog.map((log, index) => (
						<p key={`${log + index}`}>{log}</p>
					))}
				</div>
			</div>
		</main>
	);
}
