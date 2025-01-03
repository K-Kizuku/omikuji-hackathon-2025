"use client";

import { useState, useEffect } from "react";

// 型定義
type Skill = {
	id: number;
	name: string;
	attack: number;
};

type Character = {
	name: string;
	hp: number;
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
		attack: 30,
		defense: 10,
		speed: 15,
		skills: [],
	});

	const [turn, setTurn] = useState<Turn>("player");
	const [battleLog, setBattleLog] = useState<string[]>([]);

	// ダメージ計算関数
	const calculateDamage = (
		attacker: Character,
		defender: Character,
	): number => {
		const damage = Math.max(0, attacker.attack - defender.defense); // 攻撃力 - 防御力
		return damage;
	};

	// 攻撃処理
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

	// 敵のターン
	const enemyTurn = () => {
		const damage = calculateDamage(enemy, player);
		setPlayer((prev) => ({ ...prev, hp: Math.max(0, prev.hp - damage) }));
		setBattleLog((prev) => [
			...prev,
			`${enemy.name}の攻撃！ ${damage}のダメージ！`,
		]);
		setTurn("player");
	};

	// ターン管理
	useEffect(() => {
		if (turn === "enemy") {
			setTimeout(enemyTurn, 1000); // 敵のターンを1秒後に実行
		}
	}, [turn, enemyTurn]);

	// スキル使用処理
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
		<main className="flex flex-col items-center justify-between h-screen bg-gray-800 text-white">
			<h1 className="text-2xl font-bold mt-4">バトル画面</h1>

			{/* モンスターの状態 */}
			<div className="flex justify-around w-full max-w-4xl mt-8">
				{/* プレイヤー */}
				<div className="text-center">
					<p className="text-lg">{player.name}</p>
					<p>HP: {player.hp}</p>
				</div>
				{/* 敵 */}
				<div className="text-center">
					<p className="text-lg">{enemy.name}</p>
					<p>HP: {enemy.hp}</p>
				</div>
			</div>

			{/* コマンドテーブル */}
			{turn === "player" && (
				<div className="mt-8 w-full max-w-4xl">
					<table className="table-auto w-full bg-gray-700 rounded">
						<thead>
							<tr>
								<th className="px-4 py-2 border border-gray-600">コマンド</th>
								<th className="px-4 py-2 border border-gray-600">威力</th>
								<th className="px-4 py-2 border border-gray-600">行動</th>
							</tr>
						</thead>
						<tbody>
							<tr>
								<td className="px-4 py-2 border border-gray-600">攻撃</td>
								<td className="px-4 py-2 border border-gray-600">
									{player.attack}
								</td>
								<td className="px-4 py-2 border border-gray-600">
									<button
										type="button"
										onClick={handleAttack}
										className="bg-blue-500 px-4 py-2 rounded"
									>
										実行
									</button>
								</td>
							</tr>
							{player.skills.map((skill) => (
								<tr key={skill.id}>
									<td className="px-4 py-2 border border-gray-600">
										{skill.name}
									</td>
									<td className="px-4 py-2 border border-gray-600">
										{skill.attack}
									</td>
									<td className="px-4 py-2 border border-gray-600">
										<button
											type="button"
											onClick={() => handleSkill(skill)}
											className="bg-green-500 px-4 py-2 rounded"
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

			{/* バトルログ */}
			<div className="w-full max-w-4xl bg-gray-700 p-4 rounded mt-4">
				<p className="font-bold">バトルログ</p>
				<div className="h-32 overflow-y-auto">
					{battleLog.map((log, index) => (
						<p key={`$${log + index}`}>{log}</p>
					))}
				</div>
			</div>
		</main>
	);
}
