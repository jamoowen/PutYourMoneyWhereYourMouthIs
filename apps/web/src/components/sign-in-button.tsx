'use client'
import Button from "./common/button"
import SignInOptions from "./sign-in-options"

export default function SignInButton() {
  return (
    <>
      <Button
        onClick={() => {
          const dialog = document.getElementById('sign_in_modal') as HTMLDialogElement
          dialog.showModal()
        }}
        variant='outline'
        tabIndex={0}
        role="button"
      >
        <span className="font-medium text-sm">Sign In</span>
      </Button>
      <dialog id="sign_in_modal" className="modal">
        <div className={"modal-box w-[400px] absolute top-4 right-4 border bg-background border-muted shadow-lg"}>
          <SignInOptions />
        </div>
        <form method="dialog" className="modal-backdrop">
          <button className="cursor-default">close</button>
        </form>
      </dialog>

    </>

  )
}
