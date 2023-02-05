const Prompt = () => {
    let toast = (c) => {
      const { msg = "", icon = "success", position = "top-end" } = c;
      const Toast = Swal.mixin({
        toast: true,
        icon: icon,
        title: msg,
        position: position,
        showConfirmButton: false,
        timer: 3000,
        timerProgressBar: true,
        didOpen: (toast) => {
          toast.addEventListener("mouseenter", Swal.stopTimer);
          toast.addEventListener("mouseleave", Swal.resumeTimer);
        },
      });

      Toast.fire({});
    };
    let success = (c) => {
      const { msg = "", title = "", footer = "" } = c;
      Swal.fire({
        icon: "success",
        title: title,
        text: msg,
        footer: footer,
      });
    };

    let error = (c) => {
      const { msg = "", title = "", footer = "" } = c;
      Swal.fire({
        icon: "error",
        title: title,
        text: msg,
        footer: footer,
      });
    };

    return {
      toast,
      success,
      error,
      custom,
    };
  };

  async function custom(c) {
    const { icon = "", showConfirmButton = true, msg = "", title = "" } = c;
    const { value: result } = await Swal.fire({
      title: title,
      icon: icon,
      html: msg,
      focusConfirm: false,
      showCancelButton: true,
      showConfirmButton: showConfirmButton,
      willOpen: () => {
        if (c.willOpen !== undefined) {
          c.willOpen();
        }
      },
      preConfirm: () => {
        return [
          document.getElementById("start").value,
          document.getElementById("end").value,
        ];
      },
      didOpen: () => {
        if (c.didOpen !== undefined) {
          c.didOpen();
        }
      },
    });

    if (result) {
      if (result.dismiss !== Swal.DismissReason.cancel) {
        if (result.value !== "") {
          if (c.callback !== undefined) {
            c.callback(result);
          }
        } else {
          c.callback(false);
        }
      } else {
        c.callback(false);
      }
    }
  }

  