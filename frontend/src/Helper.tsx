import React, { useMemo, useRef, useState } from "react";

export function useAlert() {
  const ref = useRef<HTMLDialogElement>(null);
  const [msg,setMsg] = useState<string>()
  const dlg = <Alert ref={ref} msg={msg} />; 
  const alert = (msg?:string) => {
    setMsg(msg)
    ref.current?.showModal();
  };
  return [dlg, alert] as const;
}

const Alert = React.memo(
  React.forwardRef<HTMLDialogElement, { msg?: string }>(function (
    { msg },
    ref
  ) {
    console.debug("render Alert")
    return (
      <React.Fragment>
        <dialog ref={ref}>
          <div>{msg}</div>
          <form method="dialog" className="text-center mt-1">
            <button>OK</button>
          </form>
        </dialog>
      </React.Fragment>
    );
  }),
  (prev, next) => {
    return prev.msg === next.msg
  }
);
