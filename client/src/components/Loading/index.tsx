import Image from "next/image";

export const Loading = () => {
	return (
		<div className="min-h-screen flex flex-col items-center justify-center bg-gradient-to-br from-green-300 via-emerald-400 to-teal-500 text-white">
			<div className="relative w-32 h-32 mb-6 animate-spin-slow">
				<Image
					src="/loading.png" // アイコンのパスを指定
					alt="Loading Snake Icon"
					layout="fill"
					objectFit="contain"
				/>
			</div>
			<p className="text-xl font-bold tracking-wide animate-pulse">
				Loading...
			</p>
		</div>
	);
};
