class PaymentsController < ApplicationController
  def index
    # head :ok
    render json: Payment.all
  end
end