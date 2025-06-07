export default function Page() {
  return (
    <div className="flex flex-col items-center justify-items-center min-h-screen p-8  gap-16 sm:p-20 bg-black text-white">
      <div className="max-w-4xl text-center space-y-6">
        <h1 className="text-4xl sm:text-6xl font-bold tracking-tight">
          Put Your Money Where Your Mouth Is
        </h1>
        <p className="text-lg sm:text-xl text-gray-300">
          1v1 on rust? On the blacktop? No more trusting the other dude (who's ass you're about to whoop)
          to honor his bet. Blockchain based wagers enable a trustless system, where only a unanimous vote settles the bet. You put your money down at the beginning and only the winner withdraws.
        </p>
      </div>

      <div className="max-w-4xl space-y-10 text-center">

        <section className="space-y-4">
          <h2 className="text-2xl sm:text-3xl font-semibold">The Solution</h2>
          <p className="text-gray-400 text-base sm:text-lg">
            <strong>Put Your Money Where Your Mouth Is</strong> lets you create, accept, and lock in wagers with friends and rivals directly via wallet-to-wallet agreements. It’s trustless, transparent, and enforceable—without a middleman.
          </p>
        </section>

        <section className="space-y-4">
          <h2 className="text-2xl sm:text-3xl font-semibold">How It Works</h2>
          <ul className="text-left text-gray-300 list-disc list-inside space-y-2 max-w-xl mx-auto">
            <li>
              Create a wager and send it to your friend’s wallet address.
            </li>
            <li>
              They accept the challenge and stake the agreed-upon amount.
            </li>
            <li>
              Once all parties have staked, the bet is locked in.
            </li>
            <li>
              After the event, everyone votes on the outcome.
            </li>
            <li>
              Funds are only released if all participants vote the same: winner or cancellation.
            </li>
          </ul>
        </section>
      </div>
    </div>
  );
}
