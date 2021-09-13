FROM BASEIMAGE
RUN apk --no-cache add ca-certificates bash

ARG ARCH
ARG TINI_VERSION
ENV USER_ID=1001
# must match Makefile variable AZURERM_REFSPEC
ENV XP_TERRAFORM_PROVIDER_VERSION=2.74.0
ENV XP_TERRAFORM_PROVIDER_SOURCE=hashicorp/azurerm

# Setup Terraform environment
ENV TERRAFORM_VERSION 1.0.5
ENV PLUGIN_DIR /terraform/provider-mirror/registry.terraform.io/hashicorp/azurerm/${XP_TERRAFORM_PROVIDER_VERSION}/linux_${ARCH}
ENV TF_CLI_CONFIG_FILE /terraform/.terraformrc
ENV TF_FORK 0

RUN mkdir -p ${PLUGIN_DIR}

ADD https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_${ARCH}.zip /tmp
ADD https://releases.hashicorp.com/terraform-provider-azurerm/${XP_TERRAFORM_PROVIDER_VERSION}/terraform-provider-azurerm_${XP_TERRAFORM_PROVIDER_VERSION}_linux_${ARCH}.zip /tmp
ADD terraformrc.hcl ${TF_CLI_CONFIG_FILE}

RUN unzip /tmp/terraform_${TERRAFORM_VERSION}_linux_${ARCH}.zip -d /usr/local/bin \
  && chmod +x /usr/local/bin/terraform \
  && rm /tmp/terraform_${TERRAFORM_VERSION}_linux_${ARCH}.zip \
  && unzip /tmp/terraform-provider-azurerm_${XP_TERRAFORM_PROVIDER_VERSION}_linux_${ARCH}.zip -d ${PLUGIN_DIR} \
  && chmod +x ${PLUGIN_DIR}/* \
  && rm /tmp/terraform-provider-azurerm_${XP_TERRAFORM_PROVIDER_VERSION}_linux_${ARCH}.zip \
  && chown -R ${USER_ID}:${USER_ID} /terraform
# End of - Setup Terraform environment

ADD provider /usr/local/bin/crossplane-tf-azure-provider

USER ${USER_ID}
EXPOSE 8080
ENTRYPOINT ["crossplane-tf-azure-provider"]
