"use client";

import { type FC } from "react";
import { useFormState } from "react-dom";
import { registerAction } from "@/app/actions/register/registerAction";
import { PhoneInputMask } from "@/app/entities/phoneInputMask";
import { SubmitButton } from "@/app/entities/submitButton";
import { EFormFields } from "@/app/features/register/registerForm/enums";
import { useTranslation } from "@/app/i18n/client";
import { Input } from "@/app/uikit/components/input";
import "./RegisterForm.scss";

export const RegisterForm: FC = () => {
  const [state, formAction] = useFormState(registerAction, {});
  const { t } = useTranslation("index");

  return (
    <form action={formAction} className="Form">
      <div className="Form-FormFieldGroup">
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
      <div className="Form-Button">
        <SubmitButton buttonText={t("pages.register.button")} />
      </div>
    </form>
  );
};
