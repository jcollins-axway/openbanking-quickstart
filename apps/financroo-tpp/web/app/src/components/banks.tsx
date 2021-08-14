import apexfinancialIcon from "../assets/banks/apexfinancial-icon.svg";
import apexfinancial from "../assets/banks/apexfinancial-logo.svg";
import chaseIcon from "../assets/banks/chase-icon.svg";
import chase from "../assets/banks/chase-logo.svg";
import deutcshebank from "../assets/banks/deutcshebank-logo.svg";
import griffinBankIcon from "../assets/banks/griffinbank-icon.svg";
import griffinBank from "../assets/banks/griffinbank-logo.svg";
import hsbc from "../assets/banks/hsbc-logo.svg";
import santander from "../assets/banks/santander-logo.svg";

export type Bank = {
  value: string;
  disabled: boolean;
  name?: string;
  logo: string;
  icon?: string;
};

export const banks: Bank[] = [
  {
    value: "griffinbank",
    disabled: false,
    name: "Griffin Bank",
    logo: griffinBank,
    icon: griffinBankIcon,
  },
  {
    value: "apexfinancial",
    disabled: true,
    name: "Apex Financial",
    logo: apexfinancial,
    icon: apexfinancialIcon,
  },
  {
    value: "chase",
    disabled: true,
    logo: chase,
    icon: chaseIcon,
  },
  {
    value: "deutcshebank",
    disabled: true,
    logo: deutcshebank,
  },
  {
    value: "hsbc",
    disabled: true,
    logo: hsbc,
  },
  {
    value: "santander",
    disabled: true,
    logo: santander,
  },
];
