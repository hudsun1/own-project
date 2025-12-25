class QrsController < ApplicationController
  def index
    qrs = Qr.all
    render json: qrs
  end

  def show
    qr = Qr.find_by(id: params[:id])
    if qr
      render json: qr
    else
      render json: { error: "QR code not found" }, status: :not_found
    end
  end

  def create
    qr = Qr.new(qr_params)

    if qr.save
      render json: qr, status: :created
    else
      render json: { error: qr.errors.full_messages.join(', ') }, status: :unprocessable_entity
    end
  end

  private

  def qr_params
    # Accept both nested (qr: { imageUrl, description }) and flat (imageUrl, description) parameters
    if params[:qr].present?
      params.require(:qr).permit(:imageUrl, :description)
    else
      params.permit(:imageUrl, :description)
    end
  end
end

