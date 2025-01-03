"use client";

import { useState } from "react";
// import axios from "axios";

export default function TrainingPage() {
	const [points, setPoints] = useState(10); // 割り振り可能なポイント
	const [stats, setStats] = useState({
		hp: 0,
		attack: 0,
		defense: 0,
		speed: 0,
	});
	const [totalExperience, setTotalExperience] = useState(0); // 総経験値
	const [pythonName, setPythonName] = useState(""); // モンスター名

	// データ取得用のモック関数
	// const fetchPythonData = async () => {
	// 	// モンスターIDを仮定（通常はログイン中のユーザーやURLパラメータから取得）
	// 	const pythonId = 1;

	// 	try {
	// 		const response = await axios.get(`/api/pythons/${pythonId}`);
	// 		const { name, exp, stats: fetchedStats } = response.data;

	// 		setPythonName(name);
	// 		setTotalExperience(exp);
	// 		setStats(fetchedStats);
	// 	} catch (error) {
	// 		console.error("データの取得に失敗しました", error);
	// 	}
	// };

	// useEffect(() => {
	// 	fetchPythonData();
	// }, []);

	const handleIncrease = (stat: keyof typeof stats) => {
		if (points > 0) {
			setStats((prev) => ({ ...prev, [stat]: prev[stat] + 1 }));
			setPoints((prev) => prev - 1);
		}
	};

	const handleDecrease = (stat: keyof typeof stats) => {
		if (stats[stat] > 0) {
			setStats((prev) => ({ ...prev, [stat]: prev[stat] - 1 }));
			setPoints((prev) => prev + 1);
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
								<p>{value}</p>
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
		</main>
	);
}
