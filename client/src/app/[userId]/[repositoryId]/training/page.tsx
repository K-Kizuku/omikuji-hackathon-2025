"use client";

import { useState } from "react";

export default function TrainingPage() {
	const [points, setPoints] = useState(10);
	const [stats, setStats] = useState({
		hp: 0,
		attack: 0,
		defense: 0,
		speed: 0,
	});
	const [totalExperience, setTotalExperience] = useState(0);
	const [pythonName, setPythonName] = useState("Pythonモンスター");
	const [trainingLog, setTrainingLog] = useState<string[]>([]);
	const [bonusClaimed, setBonusClaimed] = useState(false); // ボーナス獲得済みフラグ

	// ステータス増加処理
	const handleIncrease = (stat: keyof typeof stats) => {
		if (points > 0) {
			setStats((prev) => ({ ...prev, [stat]: prev[stat] + 1 }));
			setPoints((prev) => prev - 1);
			setTrainingLog((prev) => [...prev, `${stat.toUpperCase()} +1`]);
		}
	};

	// ステータス減少処理
	const handleDecrease = (stat: keyof typeof stats) => {
		if (stats[stat] > 0) {
			setStats((prev) => ({ ...prev, [stat]: prev[stat] - 1 }));
			setPoints((prev) => prev + 1);
			setTrainingLog((prev) => [...prev, `${stat.toUpperCase()} -1`]);
		}
	};

	// ログインボーナス処理
	const claimBonus = () => {
		if (!bonusClaimed) {
			const bonusPoints = 5; // ボーナスポイント数
			setPoints((prev) => prev + bonusPoints);
			setTrainingLog((prev) => [
				...prev,
				`ログインボーナスで+${bonusPoints}ポイント獲得！`,
			]);
			setBonusClaimed(true);
		}
	};

	return (
		<main className="flex flex-col items-center justify-center h-screen bg-gray-800 text-white">
			<h1 className="text-2xl font-bold mb-4">育成画面</h1>

			{/* モンスター情報 */}
			<div className="flex flex-col items-center">
				<p className="text-xl">モンスター名: {pythonName}</p>
				<p>総経験値: {totalExperience}</p>
			</div>

			{/* ログインボーナス */}
			<div className="mt-4">
				<button
					type="button"
					className={`px-4 py-2 rounded ${
						bonusClaimed ? "bg-gray-500 cursor-not-allowed" : "bg-yellow-500"
					}`}
					onClick={claimBonus}
					disabled={bonusClaimed}
				>
					{bonusClaimed ? "ボーナス獲得済み" : "ログインボーナスを受け取る"}
				</button>
			</div>

			{/* ステータス */}
			<div className="mt-8">
				<p className="mb-4">残りポイント: {points}</p>
				<div className="grid grid-cols-4 gap-4">
					{Object.entries(stats).map(([key, value]) => {
						const translatedKey =
							key === "hp"
								? "HP"
								: key === "attack"
									? "攻撃力"
									: key === "defense"
										? "防御力"
										: "スピード";

						return (
							<div key={key} className="flex flex-col items-center">
								<p className="capitalize mb-2">{translatedKey}</p>
								<p className="text-2xl font-bold">{value}</p>
								<div className="flex space-x-2 mt-2">
									<button
										type="button"
										className="bg-green-500 px-2 py-1 rounded"
										onClick={() => handleIncrease(key as keyof typeof stats)}
										disabled={points <= 0}
									>
										+
									</button>
									<button
										type="button"
										className="bg-red-500 px-2 py-1 rounded"
										onClick={() => handleDecrease(key as keyof typeof stats)}
										disabled={stats[key as keyof typeof stats] <= 0}
									>
										-
									</button>
								</div>
							</div>
						);
					})}
				</div>
			</div>

			{/* トレーニングログ */}
			<div className="w-full max-w-4xl bg-gray-700 p-4 rounded mt-8">
				<p className="font-bold">トレーニングログ</p>
				<div className="h-32 overflow-y-auto">
					{trainingLog.map((log, index) => (
						<p key={`${log + index}`}>{log}</p>
					))}
				</div>
			</div>
		</main>
	);
}
