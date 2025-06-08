export default function Page() {
  return (
    <div className="flex flex-col items-center justify-items-center min-h-screen p-8  gap-16 sm:p-20 bg-black text-white">
      <div className="max-w-4xl text-center space-y-6">
        <h1 className="text-4xl sm:text-6xl font-bold tracking-tight">
          Put Your Money Where Your Mouth Is
        </h1>
        <p className="text-lg sm:text-xl text-gray-300">
          1v1 on rust? First to 11 On the blacktop? Saturday morning matchplay?
          We've all been in the situation where we bet on the game's we play.
          We have to trust the other player to honor their side of the bet (In the event you whoop their ass).
          BUT sometimes they dont honor their bet...
          I for won have not been paid many a time for wagers I have fairly won
        </p>
      </div>

      <div className="max-w-4xl space-y-10 text-center">

        <section className="space-y-4">
          <h2 className="text-2xl sm:text-3xl font-semibold">The Solution</h2>
          <p className="text-gray-400 text-base sm:text-lg">
            <strong>Put Your Money Where Your Mouth Is</strong> lets you create, accept, and lock in wagers with friends and rivals using the blockchain to protect and release the wagered funds. It’s trustless, transparent, and prevents dishonest behavior.
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
              Funds are only released if a unanimous vote is reached.
            </li>
            <li className="opacity-50 text-sm">
              You can keep revoting if you change your mind on the outcome.
              The absolute worst case scenario is that a unanimous decision is never reached
              and so the money will remain locked until all players agree on the outcome.
            </li>
          </ul>
        </section>
      </div>
    </div>
  );
}
