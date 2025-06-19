

export default function Page() {
  return (
    <div className="flex flex-col space-y-6">
      <div className="max-w-4xl text-center space-y-6">
        <h1 className="text-4xl sm:text-6xl font-bold tracking-tight">
          About
        </h1>
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
  )
}
