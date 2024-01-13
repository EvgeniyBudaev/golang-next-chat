"use client";

import isNil from "lodash/isNil";
import { redirect } from "next/navigation";
import { type FC, useEffect } from "react";
import { useFormState } from "react-dom";
import { registerAction } from "@/app/actions/register/registerAction";
import { PhoneInputMask } from "@/app/entities/phoneInputMask";
import { SubmitButton } from "@/app/entities/submitButton";
import { EFormFields } from "@/app/features/register/registerForm/enums";
import { useTranslation } from "@/app/i18n/client";
import { ERoutes } from "@/app/shared/enums";
import { createPath } from "@/app/shared/utils";
import { Input } from "@/app/uikit/components/input";
import { notify } from "@/app/uikit/components/toast/utils";
import "./RegisterForm.scss";

export const RegisterForm: FC = () => {
  const initialState = {
    data: undefined,
    error: undefined,
    errors: undefined,
    success: false,
  };
  const [state, formAction] = useFormState(registerAction, initialState);
  const { t } = useTranslation("index");
  console.log("RegisterForm state: ", state);

  useEffect(() => {
    if (state?.error) {
      notify.error({ title: state?.error });
    }
    if (!isNil(state.data) && state.success && !state?.error) {
      notify.success({ title: "Ok" });
      return redirect(
        createPath({
          route: ERoutes.Login,
        }),
      );
    }
  }, [state]);

  return (
    <form action={formAction} className="RegisterForm">
      <div className="RegisterForm-FieldGroup">
        <Input
          errors={state?.errors?.userName}
          isRequired={true}
          label={t("form.userName") ?? "User name"}
          name={EFormFields.UserName}
          type="text"
        />
        <Input
          errors={state?.errors?.firstName}
          isRequired={true}
          label={t("form.firstName") ?? "First Name"}
          name={EFormFields.FirstName}
          type="text"
        />
        <Input
          errors={state?.errors?.lastName}
          isRequired={true}
          label={t("form.lastName") ?? "Last Name"}
          name={EFormFields.LastName}
          type="text"
        />
        <PhoneInputMask
          errors={state?.errors?.mobileNumber}
          isRequired={true}
          label={t("form.mobileNumber") ?? "Mobile phone"}
          name={EFormFields.MobileNumber}
        />
        <Input
          errors={state?.errors?.email}
          isRequired={true}
          label={t("form.email") ?? "Email"}
          name={EFormFields.Email}
          type="text"
        />
        <Input
          errors={state?.errors?.password}
          isRequired={true}
          label={t("form.password") ?? "Password"}
          name={EFormFields.Password}
          type="text"
        />
        <Input
          errors={state?.errors?.passwordConfirm}
          isRequired={true}
          label={t("form.passwordConfirm") ?? "Password confirm"}
          name={EFormFields.PasswordConfirm}
          type="text"
        />
      </div>
      <div className="RegisterForm-Button">
        <SubmitButton buttonText={t("pages.register.button")} />
      </div>
    </form>
  );
};
