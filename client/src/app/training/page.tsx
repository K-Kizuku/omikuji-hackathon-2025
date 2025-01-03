"use client";

import { useState } from "react";

export default function TrainingPage() {
	const [points, setPoints] = useState(10); // 自由に割り振れるポイント数
	const [stats, setStats] = useState({
		attack: 0,
		speed: 0,
		health: 0,
	});

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

			{/* パイモンの画像 */}
			<div className="flex flex-col items-center">
				<img src="monster-image-url" alt="Monster" className="w-32 h-32 mb-2" />
				<p>パイモン</p>
			</div>

			<div className="mt-8">
				<p className="mb-4">Remaining Points: {points}</p>
				<div className="grid grid-cols-3 gap-4">
					{Object.entries(stats).map(([key, value]) => (
						<div key={key} className="flex flex-col items-center">
							<p className="capitalize mb-2">{key}</p>
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
					))}
				</div>
			</div>
		</main>
	);
}
