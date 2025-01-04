"use client";

import Link from "next/link";
import { useState } from "react";

// 型定義
type PymonStats = {
	hp: number;
	attack: number;
	defense: number;
	speed: number;
};

type Pymon = {
	id: string;
	name: string;
	exp: number;
	stats: PymonStats;
};

export default function HomePage() {
	const [selectedPymon, setSelectedPymon] = useState<Pymon | null>(null); // 選択されたパイモン

	// ダミーデータ
	const pymons: Pymon[] = [
		{
			id: "1",
			name: "Pyro",
			exp: 1200,
			stats: {
				hp: 100,
				attack: 50,
				defense: 30,
				speed: 20,
			},
		},
		{
			id: "2",
			name: "Hydra",
			exp: 800,
			stats: {
				hp: 120,
				attack: 40,
				defense: 40,
				speed: 15,
			},
		},
		{
			id: "3",
			name: "Electro",
			exp: 1500,
			stats: {
				hp: 80,
				attack: 60,
				defense: 20,
				speed: 25,
			},
		},
	];

	// 詳細パネルを開く
	const openDetails = (pymon: Pymon) => {
		setSelectedPymon(pymon);
	};

	// 詳細パネルを閉じる
	const closeDetails = () => {
		setSelectedPymon(null);
	};

	// ログアウトボタンの動作（仮）
	const handleLogout = () => {
		alert("ログアウトしました！");
	};

	return (
		<div className="min-h-screen bg-gradient-to-br from-blue-600 via-blue-400 to-yellow-400 text-white p-8 relative">
			<header className="flex justify-between items-center bg-gray-900 bg-opacity-80 p-4 rounded-lg">
				<h1 className="text-3xl font-mono font-bold tracking-wide">
					Python World
				</h1>
				<div className="flex space-x-4">
					<button
						type="button"
						onClick={handleLogout}
						className="bg-red-500 text-white py-2 px-4 rounded-lg hover:bg-red-700 transition-transform transform hover:scale-105"
					>
						ログアウト
					</button>
					<Link
						href="/create"
						className="bg-green-500 text-white py-2 px-4 rounded-lg hover:bg-green-700 transition-transform transform hover:scale-105"
					>
						ペイモン作成
					</Link>
				</div>
			</header>
			<main className="mt-8">
				<h2 className="text-2xl font-bold mb-4">Your Pymons</h2>
				<div className="grid grid-cols-1 md:grid-cols-3 gap-8">
					{pymons.map((pymon) => (
						<div
							key={pymon.id}
							className="bg-gray-800 rounded-lg p-4 shadow-lg text-center"
						>
							<h3 className="text-lg font-bold">{pymon.name}</h3>
							<p>Exp: {pymon.exp}</p>
							<p>HP: {pymon.stats.hp}</p>
							<p>Attack: {pymon.stats.attack}</p>
							<div className="mt-4 flex justify-center space-x-4">
								<Link
									href={`/${pymon.id}/育成`}
									className="bg-green-500 text-white py-2 px-4 rounded hover:bg-green-700"
								>
									育成
								</Link>
								<Link
									href={`/${pymon.id}/戦闘`}
									className="bg-red-500 text-white py-2 px-4 rounded hover:bg-red-700"
								>
									戦闘
								</Link>
							</div>
							<button
								type="button"
								onClick={() => openDetails(pymon)}
								className="mt-4 bg-blue-500 text-white py-2 px-4 rounded hover:bg-blue-700"
							>
								詳細を見る
							</button>
						</div>
					))}
				</div>
			</main>

			{/* サイドパネル */}
			<div
				className={`fixed top-0 right-0 h-full w-1/3 bg-gray-900 bg-opacity-90 p-6 shadow-lg z-50 transform transition-transform ${
					selectedPymon ? "translate-x-0" : "translate-x-full"
				}`}
			>
				{selectedPymon && (
					<>
						<button
							type="button"
							onClick={closeDetails}
							className="absolute top-4 right-4 text-white text-2xl font-bold"
						>
							&times;
						</button>
						<h3 className="text-2xl font-bold mb-4">{selectedPymon.name}</h3>
						<p>Exp: {selectedPymon.exp}</p>
						<p>HP: {selectedPymon.stats.hp}</p>
						<p>Attack: {selectedPymon.stats.attack}</p>
						<p>Defense: {selectedPymon.stats.defense}</p>
						<p>Speed: {selectedPymon.stats.speed}</p>
					</>
				)}
			</div>
		</div>
	);
}
